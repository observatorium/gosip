package tenant

import "net/url"

//Static database of tenants
type Static struct {
	tenants map[string]Tenant
}

//NewStatic returns a static database with given tenants
func NewStatic(tenants map[string]Tenant) *Static {
	return &Static{tenants: tenants}
}

//PrometheusURL returns the tenant's Prometheus URL by its username
func (s *Static) PrometheusURL(username string) (*url.URL, error) {
	t, found := s.tenants[username]
	if !found {
		return nil, ErrTenantNotFound
	}

	return t.PrometheusURL, nil
}
