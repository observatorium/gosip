package users

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{db: db}
}

type postgresUser struct {
	ID       string
	username string
	token    string
	password string
}

func (u postgresUser) Username() string {
	return u.username
}

func (u postgresUser) Token() string {
	return u.token
}

func (u postgresUser) Password() string {
	return u.password
}

func (p *Postgres) GetByUsername(username string) (User, error) {
	row := p.db.QueryRowContext(context.TODO(),
		`SELECT id, username, token, password FROM users WHERE username = $1 LIMIT 1`,
		username,
	)

	var user postgresUser
	err := row.Scan(&user.ID, &user.username, &user.token, &user.password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, errors.Wrap(err, "failed to get user by username")
	}

	return user, nil
}

func (p *Postgres) GetByToken(token string) (User, error) {
	row := p.db.QueryRowContext(context.TODO(),
		`SELECT id, username, token, password FROM users WHERE token = $1`,
		token,
	)

	var user postgresUser
	err := row.Scan(&user.ID, &user.username, &user.token, &user.password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, errors.Wrap(err, "failed to get user by token")
	}

	return user, nil
}
