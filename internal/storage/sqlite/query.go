package sqlite

const (
	saveUser = "INSERT INTO users(email, pass_hash) VALUES (?, ?)"
	getUser  = "SELECT id, email, pass_hash FROM users WHERE email=?"
	isAdmin  = "SELECT is_admin FROM users WHERE id=?"
	getApp   = "SELECT id, name, secret FROM app WHERE id = ?"
)
