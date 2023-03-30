package model

type User struct {
	UserName     string
	PasswordHash string

	CreatedAt int64
	UpdatedAt int64
}

type UserInfo struct {
	UserName string
}
