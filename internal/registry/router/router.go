package router

import (
	"github.com/gorilla/mux"
	"github.com/patriuk/hatch/internal/registry/discovery"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/register", discovery.Register)
	r.HandleFunc("/unregister", discovery.Unregister)
	r.HandleFunc("/refresh", discovery.Refresh)

	return r
}
