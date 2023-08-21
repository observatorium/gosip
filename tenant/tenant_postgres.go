package tenant

import (
	"context"
	"database/sql"
	"net/url"

	"github.com/pkg/errors"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{db: db}
}

func (p *Postgres) PrometheusURL(username string) (*url.URL, error) {
	query := `
SELECT tenants.prometheus FROM tenants
LEFT JOIN tenants_users tu ON tenants.id = tu.tenant_id
LEFT JOIN users ON tu.user_id = users.id
WHERE users.username = $1
`
	row := p.db.QueryRowContext(context.TODO(), query, username)

	var prometheusURL string
	if err := row.Scan(&prometheusURL); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrTenantNotFound
		}

		return nil, errors.Wrap(err, "failed to scan prometheus URL")
	}

	return url.Parse(prometheusURL)
}
