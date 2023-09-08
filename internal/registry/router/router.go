package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/patriuk/hatch/internal/registry/handlers"
)

func SetupRoutes(h handlers.Handlers) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/register", h.Discovery.Register).
		Methods(http.MethodPut)

	r.HandleFunc("/unregister", h.Discovery.Unregister).
		Methods(http.MethodDelete)

	r.HandleFunc("/refresh", h.Discovery.Refresh).
		Methods(http.MethodPost)

	r.HandleFunc("/services", h.Discovery.GetServices).
		Methods(http.MethodGet)

	// todo add GetAllByName with param name

	r.HandleFunc("/test", h.Discovery.Test).
		Methods(http.MethodGet)

	return r
}
