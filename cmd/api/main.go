package main

import (
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":0")

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello darkness my old friend \n"))
	})

	srv := &http.Server{
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Starting server on %s", listener.Addr())

	err = srv.Serve(listener)
	log.Fatal(err)
}
