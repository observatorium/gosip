package users

import "errors"

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrPasswordIncorrect = errors.New("password incorrect")
)

type UserDatabase interface {
	GetByUsername(username string) (User, error)
	GetByToken(token string) (User, error)
}

type User interface {
	Name() string
	Token() string
	Password() string
}
