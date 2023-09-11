package handlers

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/patriuk/hatch/internal/gateway/balancer"
	"github.com/patriuk/hatch/internal/gateway/repositories"
)

type GatewayHandler struct {
	Balancer balancer.Balancer
	repo     repositories.ServiceRepository
}

func NewGatewayHandler(balancer balancer.Balancer, repo repositories.ServiceRepository) *GatewayHandler {
	return &GatewayHandler{
		Balancer: balancer,
		repo:     repo,
	}
}

func (h *GatewayHandler) Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ping")
}

func (h *GatewayHandler) RouteRequest(w http.ResponseWriter, r *http.Request) {
	parts := strings.SplitN(r.URL.String()[1:], "/", 2)
	serviceName := parts[0]
	fmt.Println("parts1 empty? --", parts[1]) // actually panics when it has nothing
	// to split.. rip
	rest := "/" + parts[1]
	fmt.Println("rest --", rest)

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
