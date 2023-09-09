package discovery

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/patriuk/hatch/internal/common"
	"github.com/patriuk/hatch/internal/helpers"
	"github.com/patriuk/hatch/internal/registry/repositories/service"
)

type Handler struct {
	repo service.ServiceRepository
}

func NewHandler(repo service.ServiceRepository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("registry service - register handler")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	s := &common.Service{
		Timestamp: time.Now().Unix(),
	}
	err = json.Unmarshal(body, s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(helpers.PrettyPrint(s))

	h.repo.Register(*s)
}

func (h *Handler) Unregister(w http.ResponseWriter, r *http.Request) {
	fmt.Println("registry service - unregister handler")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	s := &common.Service{}
	err = json.Unmarshal(body, s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(helpers.PrettyPrint(s))

	h.repo.Unregister(*s)
}

func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request) {
	fmt.Println("registry service - refresh handler")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	s := &common.Service{
		Timestamp: time.Now().Unix(),
	}
	err = json.Unmarshal(body, s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(helpers.PrettyPrint(s))

	h.repo.Refresh(*s)
}
