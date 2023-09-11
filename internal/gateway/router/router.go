package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/patriuk/hatch/internal/gateway/handlers"
)

func SetupRoutes(h *handlers.GatewayHandler) *mux.Router {
	r := mux.NewRouter()

	r.PathPrefix("/").HandlerFunc(h.RouteRequest)

	r.HandleFunc("/ping", h.Ping).Methods(http.MethodGet)

	return r
}
