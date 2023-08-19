package registry

import (
	"net"
	"net/http"

	"github.com/patriuk/hatch/internal/registry/config"
	"github.com/patriuk/hatch/internal/registry/router"
	"github.com/patriuk/hatch/internal/registry/server"
)

type Registry struct {
	Config   config.Config
	listener net.Listener
	router   http.Handler
}

type Params struct {
	Config   config.Config
	Listener net.Listener
}

func New(params Params) *Registry {
	registry := &Registry{
		Config:   params.Config,
		listener: params.Listener,
		router:   router.NewRouter(),
	}

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

	return registry
}

//	func (app *App) setupRoutes() {
//	    // Initialize and add your API handlers to the router
//	    apiHandler := NewAPIHandler(app.db)
//	    app.router.HandleFunc("/api/resource", apiHandler.HandleResource).Methods("GET")
//	    // Add more routes as needed
//	}
//
//	func (app *App) setupMiddleware() {
//	    // Add your middleware to the router, if any
//	    // Example: app.router.Use(middleware.MyMiddleware)
//	}

func (registry *Registry) Serve() error {
	return server.Serve(registry.listener, registry.router)
}
