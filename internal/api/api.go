package api

import (
	"net"

	"github.com/patriuk/hatch/internal/api/config"
	"github.com/patriuk/hatch/internal/api/server"
)

type Api struct {
	Config   config.Config
	listener net.Listener
}

type Params struct {
	Config   config.Config
	Listener net.Listener
}

// TODO: use a standard lib router for service
// TODO: use gorilla mux router for registry
// TODO: figure out how to handle routing in gateway. do we even need a complex
// router given we're just proxying to other services?

func New(params Params) *Api {
	api := &Api{
		Config:   params.Config,
		listener: params.Listener,
	}

	// boilerplate
	// app := &App{
	// 	router: mux.NewRouter(),
	// }
	//
	// // Initialize and connect to the database
	// db, err := NewDatabase()
	// if err != nil {
	// 	// Handle the error
	// }
	// app.db = db
	//
	// // Wire up your routes and middleware
	// app.setupRoutes()
	// app.setupMiddleware()

	return api
}

func (api *Api) Serve() error {
	return server.Serve(api.listener)
}
