package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func ListenAndServe(ip string, port uint16, router http.Handler) error {
	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", ip, port),
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Starting server on %s", srv.Addr)

	err := srv.ListenAndServe()
	log.Fatal(err)

	return nil
}
