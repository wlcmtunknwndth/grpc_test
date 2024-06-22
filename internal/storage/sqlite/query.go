package sqlite

const (
	saveUser = "INSERT INTO users(email, pass_hash) VALUES (?, ?)"
)
