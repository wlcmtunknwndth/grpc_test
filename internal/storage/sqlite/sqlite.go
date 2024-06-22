package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"
	"github.com/wlcmtunknwndth/grpc_test/internal/domain/models"
	"github.com/wlcmtunknwndth/grpc_test/internal/storage"
)

type Storage struct {
	db *sql.DB
}

const scope = "internal.storage.sqlite.sqlite."

func New(storagePath string) (*Storage, error) {
	const op = scope + "New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) SaveUser(ctx context.Context, email string, passHash []byte) (int64, error) {
	const op = scope + "SaveUser"

	stmt, err := s.db.Prepare(saveUser)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	res, err := stmt.Exec(ctx, email, passHash)
	if err != nil {
		var sqliteErr sqlite3.Error

		if errors.As(err, &sqliteErr) && errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
			return 0, fmt.Errorf("%s: %w", op, storage.ErrUserExits)
		}

		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (s *Storage) User(ctx context.Context, email string) (models.User, error) {
	const op = scope + "User"

	stmt, err := s.db.Prepare(getUser)
	if err != nil {
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	res := stmt.QueryRowContext(ctx, email)

	usr := models.User{}

	if err = res.Scan(&usr.ID, &usr.Email, &usr.PassHash); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, fmt.Errorf("%s: %w", op, storage.ErrUserNotFound)
		}

		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return usr, nil
}

func (s *Storage) IsAdmin(ctx context.Context, id int64) (bool, error) {
	const op = scope + "IsAdmin"

	stmt, err := s.db.Prepare(isAdmin)
	if err != nil {
		return false, fmt.Errorf("%s: %w", op, err)
	}

	res := stmt.QueryRowContext(ctx, id)

	var ans bool

	if err = res.Scan(&ans); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, fmt.Errorf("%s: %w", op, storage.ErrAppNotFound)
		}

		return false, fmt.Errorf("%s: %w", op, err)
	}

	return ans, nil
}

func (s *Storage) App(ctx context.Context, appID int) (models.App, error) {
	const op = scope + "App"

	stmt, err := s.db.Prepare(getApp)
	if err != nil {
		return models.App{}, fmt.Errorf("%s: %w", op, err)
	}

	var app models.App

	res := stmt.QueryRowContext(ctx, &app.ID, &app.Name, &app.Secret)

	if err = res.Scan(&app); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.App{}, fmt.Errorf("%s: %w", op, storage.ErrAppNotFound)
		}
		return models.App{}, fmt.Errorf("%s: %w", op, err)
	}

	return app, nil
}
