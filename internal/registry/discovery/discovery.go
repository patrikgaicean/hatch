package discovery

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/patriuk/hatch/internal/helpers"
)

type Discovery struct {
	Name     string `json:"name"`
	IP       string `json:"ip"`
	Port     uint16 `json:"port"`
	Protocol string `json:"protocol"`
	IPType   string `json:"ipType"`
	// Address  string `json:"address"` -- add in redis though
}

// todo: validation
func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("registry service - register handler")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	v := &Discovery{}
	err = json.Unmarshal(body, v)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(helpers.PrettyPrint(v))
	// logic to register -> add in redis
}

func Unregister(w http.ResponseWriter, r *http.Request) {
	fmt.Println("registry service - unregister handler")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	// probably need to create a hash of details to keep as key
	v := &Discovery{}
	err = json.Unmarshal(body, v)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(helpers.PrettyPrint(v))
	// logic to unregister -> delete from redis
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	fmt.Println("registry service - refresh handler")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	// probably need to create a hash of details to keep as key
	v := &Discovery{}
	err = json.Unmarshal(body, v)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(helpers.PrettyPrint(v))
	// logic to refresh -> update timestamp in redis?
}

// probably need to implement a routine to clean the registry (redis) every
// whatever seconds/time interval. so having a timestamp for the data would
// probably be useful..
