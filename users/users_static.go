package users

import (
	"github.com/Go-SIP/gosip/config"
)

type StaticUsersDatabase struct {
	configuredUsers []*config.User
	nameIndex       map[string]*config.User
	tokenIndex      map[string]*config.User
}

type StaticUserEntry struct {
	user *config.User
}

func NewStaticUsersDatabase(configuredUsers []*config.User) UserDatabase {
	nameIndex := make(map[string]*config.User, len(configuredUsers))
	tokenIndex := make(map[string]*config.User, len(configuredUsers))

	for _, u := range configuredUsers {
		nameIndex[u.Username] = u
		tokenIndex[u.Token] = u
	}

	return &StaticUsersDatabase{
		configuredUsers: configuredUsers,
		nameIndex:       nameIndex,
		tokenIndex:      tokenIndex,
	}
}

func (d *StaticUsersDatabase) GetByUsername(username string) (User, error) {
	u, found := d.nameIndex[username]
	if !found {
		return nil, ErrUserNotFound
	}

	return &StaticUserEntry{user: u}, nil
}

func (d *StaticUsersDatabase) GetByToken(token string) (User, error) {
	u, found := d.tokenIndex[token]
	if !found {
		return nil, ErrUserNotFound
	}

	return &StaticUserEntry{user: u}, nil
}

func (u *StaticUserEntry) Name() string {
	return u.user.Username
}

func (u *StaticUserEntry) Token() string {
	return u.user.Token
}

func (u *StaticUserEntry) Password() string {
	return u.user.Password
}
