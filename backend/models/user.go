package models

type UserID int

type User struct {
	ID           UserID
	Username     string
	PasswordHash string
}

func NewUser(username string, passwordHash string) User {
	return User{
		Username:     username,
		PasswordHash: passwordHash,
	}
}
