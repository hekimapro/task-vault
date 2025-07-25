package queries

var ReadUserByEmail = "SELECT * FROM users WHERE email = $1"