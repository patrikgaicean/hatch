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

	return r
}
