package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prateek-hegde/go-social/cmd/api/routes"
	"github.com/prateek-hegde/go-social/internal/store"
)

type application struct {
	config config
	store  store.Storage
}

type config struct {
	addr string
	db   dbConfig
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdelConn  int
	maxIdleTime  string
}

func (app *application) run() error {
	// mux := http.NewServeMux()
	server := gin.Default()

	routes.GenerateRoutes(server)

	return server.Run(app.config.addr)
}
