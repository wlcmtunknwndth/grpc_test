-- CHANGE is_admin column to individual table, because evry
-- usrs mostly hold the unnecessary information, when there are
-- few admins

ALTER TABLE users
    ADD COLUMN is_admin BOOLEAN NOT NULL DEFAULT FALSE;