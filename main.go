package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/Go-SIP/gosip/auth"
	"github.com/Go-SIP/gosip/config"
	"github.com/Go-SIP/gosip/ui"
	"github.com/Go-SIP/gosip/users"
)

func main() {
	os.Exit(Main())
}

func Main() int {
	c, err := config.LoadFile("config.yaml")
	if err != nil {
		fmt.Println("Failed to load config:", err)
		return 1
	}
	promURL, err := url.Parse(c.Prometheus.URL)
	if err != nil {
		fmt.Println("Failed to parse Prometheus URL:", err)
		return 1
	}

	jaegerURL, err := url.Parse(c.Jaeger.URL)
	if err != nil {
		fmt.Println("Failed to parse Jaeger URL:", err)
		return 1
	}

	mux := http.NewServeMux()

	db := users.NewStaticUsersDatabase(c.Users)
	a := auth.NewAuthHandler(db)
	ui := ui.New()

	mux.Handle("/prometheus", a.TokenAuth(httputil.NewSingleHostReverseProxy(promURL)))
	mux.Handle("/jaeger", a.TokenAuth(httputil.NewSingleHostReverseProxy(jaegerURL)))
	mux.Handle("/", a.BasicAuth(ui))

	log.Fatal(http.ListenAndServe(":8080", mux))

	return 0
}
