package config

import (
	"io/ioutil"

	"github.com/ghodss/yaml"
)

type Config struct {
	Prometheus PrometheusConfig `json:"prometheus"`
	Jaeger     JaegerConfig     `json:"jaeger"`
	Tenants    []*Tenant        `json:"tenants"`
	Users      []*User          `json:"users"`
}

type PrometheusConfig struct {
	URL string `json:"url"`
}

type JaegerConfig struct {
	URL string `json:"url"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type Tenant struct {
	ID    string   `json:"id"`
	Users []string `json:"users"`
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
