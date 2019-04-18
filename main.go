package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/Go-SIP/gosip/auth"
	"github.com/Go-SIP/gosip/config"
	"github.com/Go-SIP/gosip/proxy"
	"github.com/Go-SIP/gosip/tenant"
	"github.com/Go-SIP/gosip/ui"
	"github.com/Go-SIP/gosip/users"

	_ "github.com/lib/pq"
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

	db, err := sql.Open("postgres", c.Database.DSN)
	if err != nil {
		fmt.Println("failed to open database:", err)
		return 1
	}

	tenants := tenant.NewPostgres(db)
	users := users.NewPostgres(db)
	auth := auth.NewHandler(users)

	mux := http.NewServeMux()
	mux.Handle("/prometheus/", auth.Token(proxy.NewPrometheus(tenants)))
	//mux.Handle("/jaeger", auth.Token(httputil.NewSingleHostReverseProxy(jaegerURL)))
	mux.Handle("/", auth.Basic(ui.New()))

	fmt.Println("Running server on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println("Failed to run server:", err)
		return 2
	}

	return 0
}
