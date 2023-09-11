package handlers

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/patriuk/hatch/internal/gateway/balancer"
	"github.com/patriuk/hatch/internal/gateway/manager"
	"github.com/patriuk/hatch/internal/gateway/repositories"
)

type GatewayHandler struct {
	Balancer balancer.Balancer
	repo     repositories.ServiceRepository
	manager  *manager.Manager
}

func NewGatewayHandler(balancer balancer.Balancer, repo repositories.ServiceRepository, manager *manager.Manager) *GatewayHandler {
	return &GatewayHandler{
		Balancer: balancer,
		repo:     repo,
		manager:  manager,
	}
}

func (h *GatewayHandler) Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ping")
}

func (h *GatewayHandler) RouteRequest(w http.ResponseWriter, r *http.Request) {
	serviceName, rest := getUrlParts(r.URL.String())
	ok := h.manager.HasService(serviceName)
	if !ok {
		http.Error(w, "Invalid service", http.StatusBadRequest)
		return
	}

	instance, err := h.Balancer.GetNextInstance(serviceName)
	if err != nil {
		http.Error(w, "Service not available", http.StatusServiceUnavailable)
		return
	}

	targetURL, err := url.Parse(fmt.Sprintf(
		"http://%s:%d%s",
		instance.IP,
		instance.Port,
		rest,
	))
	if err != nil {
		http.Error(w, "Bad target URL", http.StatusBadRequest)
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	r.URL.Host = targetURL.Host
	r.URL.Scheme = targetURL.Scheme
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
	r.Host = targetURL.Host

	proxy.ServeHTTP(w, r)
}

func getUrlParts(path string) (string, string) {
	if strings.Index(path, "/") == 0 {
		path = path[1:]
	}

	parts := strings.SplitN(path, "/", 2)

	serviceName := parts[0]
	rest := "/"

	if len(parts) == 2 {
		rest += parts[1]
	}

	return serviceName, rest
}
