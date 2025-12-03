package models

// User Register业务logic中使用
type User struct {
	UserID   int64  `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
	IsAdmin  bool   `db:"is_admin"`
	Token    string
}
