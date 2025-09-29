package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prateek-hegde/go-social/cmd/api/routes"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) run() error {
	// mux := http.NewServeMux()
	server := gin.Default()

	routes.GenerateRoutes(server)

	return server.Run(app.config.addr)
}
