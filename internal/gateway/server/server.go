package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// refactor after creating router module
// func ListenAndServe(ip string, port uint16, router http.Handler) error {
func ListenAndServe(ip string, port uint16) error {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello darkness my old friend from the gateway\n"))
	})

	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", ip, port),
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Starting server on %s", srv.Addr)

	err := srv.ListenAndServe()
	log.Fatal(err)

	return nil
}
