package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/patriuk/hatch/internal/registry/discovery"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/register", discovery.Register).
		Methods(http.MethodPut)

	r.HandleFunc("/unregister", discovery.Unregister).
		Methods(http.MethodDelete)

	r.HandleFunc("/refresh", discovery.Refresh).
		Methods(http.MethodPost)

	r.HandleFunc("/services", discovery.GetServices).
		Methods(http.MethodGet)

	return r
}
