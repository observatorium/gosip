package config

import (
	"io/ioutil"
	"net/url"

	"github.com/Go-SIP/gosip/tenant"
	"github.com/ghodss/yaml"
	"github.com/pkg/errors"
)

type Config struct {
	Tenants []*Tenant `json:"tenants"`
	Users   []*User   `json:"users"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type Tenant struct {
	ID         string           `json:"id"`
	Users      []string         `json:"users"`
	Prometheus PrometheusConfig `json:"prometheus"`
	Jaeger     JaegerConfig     `json:"jaeger"`
}

type PrometheusConfig struct {
	URL string `json:"url"`
}

type JaegerConfig struct {
	URL string `json:"url"`
}

func LoadFile(filename string) (*Config, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	c := Config{}
	err = yaml.Unmarshal(content, &c)
	return &c, err
}

//Tenants returns a map of username and tenant as read from the config
func Tenants(config *Config) (map[string]tenant.Tenant, error) {
	tenants := make(map[string]tenant.Tenant)

	for _, t := range config.Tenants {
		for _, u := range t.Users {
			jaegerURL, err := url.Parse(t.Jaeger.URL)
			if err != nil {
				return nil, errors.Wrap(err, "failed to parse Jaeger URL")
			}

			prometheusURL, err := url.Parse(t.Prometheus.URL)
			if err != nil {
				return nil, errors.Wrap(err, "failed to parse Prometheus URL")
			}

			tenants[u] = tenant.Tenant{
				ID:            t.ID,
				JaegerURL:     jaegerURL,
				PrometheusURL: prometheusURL,
			}
		}
	}

	return tenants, nil
}
