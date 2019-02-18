package users

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
)

type User interface {
	Name() string
	Token() string
	Password() string
}
