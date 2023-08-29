package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/patriuk/hatch/internal/registry/handlers"
)

func SetupRoutes(handlers handlers.Handlers) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/register", handlers.Discovery.Register).
		Methods(http.MethodPut)

	r.HandleFunc("/unregister", handlers.Discovery.Unregister).
		Methods(http.MethodDelete)

	r.HandleFunc("/refresh", handlers.Discovery.Refresh).
		Methods(http.MethodPost)

	r.HandleFunc("/services", handlers.Discovery.GetServices).
		Methods(http.MethodGet)

	return r
}
