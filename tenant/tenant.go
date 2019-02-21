package tenant

import (
	"net/url"

	"github.com/pkg/errors"
)

var (
	//ErrTenantNotFound if the tenant could not be found by ID or username
	ErrTenantNotFound = errors.New("tenant not found")
)

//Tenant of gosip
type Tenant struct {
	ID            string
	JaegerURL     *url.URL
	PrometheusURL *url.URL
}
