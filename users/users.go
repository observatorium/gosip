package users

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
)

type User interface {
	Username() string
	Token() string
	Password() string
}
