package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", CommonHandler)
	r.HandleFunc("/first", FirstHandler)
	r.HandleFunc("/second", SecondHandler)

	return r
}

func CommonHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("message from common handler")
}

func FirstHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("message from first handler")
}

func SecondHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("message from second handler")
}
