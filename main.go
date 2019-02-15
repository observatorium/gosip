package main

import (
	"fmt"
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
	auth := auth.NewHandler(db)
	ui := ui.New()

	mux.Handle("/prometheus", auth.Token(httputil.NewSingleHostReverseProxy(promURL)))
	mux.Handle("/jaeger", auth.Token(httputil.NewSingleHostReverseProxy(jaegerURL)))
	mux.Handle("/", auth.Basic(ui))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Failed to run server:", err)
		return 2
	}

	return 0
}
