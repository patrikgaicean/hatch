package discovery

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/patriuk/hatch/internal/helpers"
	"github.com/patriuk/hatch/internal/registry/repositories/service"
)

type Handler struct {
	repo service.ServiceRepository
}

func NewHandler(repo service.ServiceRepository) *Handler {
	return &Handler{
		repo: repo,
	}
}

// todo: validation
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("registry service - register handler")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	s := &service.Service{}
	err = json.Unmarshal(body, s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(helpers.PrettyPrint(s))

	h.repo.Register(service.Service{
		Name:     s.Name,
		IP:       s.IP,
		Port:     s.Port,
		Protocol: s.Protocol,
		IPType:   s.IPType,
	})
}

func (h *Handler) Unregister(w http.ResponseWriter, r *http.Request) {
	fmt.Println("registry service - unregister handler")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	s := &service.Service{}
	err = json.Unmarshal(body, s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(helpers.PrettyPrint(s))

	h.repo.Unregister(service.Service{
		Name:     s.Name,
		IP:       s.IP,
		Port:     s.Port,
		Protocol: s.Protocol,
		IPType:   s.IPType,
	})
}

func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request) {
	fmt.Println("registry service - refresh handler")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	s := &service.Service{}
	err = json.Unmarshal(body, s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(helpers.PrettyPrint(s))

	h.repo.Refresh(service.Service{
		Name:      s.Name,
		IP:        s.IP,
		Port:      s.Port,
		Protocol:  s.Protocol,
		IPType:    s.IPType,
		Timestamp: time.Now().Unix(),
	})
}

func (h *Handler) GetServices(w http.ResponseWriter, r *http.Request) {
	h.repo.GetAllByName("something")
}

// probably need to implement a routine to clean the registry (redis) every
// whatever seconds/time interval. so having a timestamp for the data would
// probably be useful..
