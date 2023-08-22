package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

type config struct {
	name string
	env  string
	port int
}

type application struct {
	config config
}

func main() {
	ip := flag.String("ip", "127.0.0.1", "Gateway ip")
	port := flag.Int("port", 8000, "Gateway port")
	flag.Parse()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello darkness my old friend from the gateway\n"))
	})

	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", *ip, *port),
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Starting server on %s", srv.Addr)

	err := srv.ListenAndServe()
	log.Fatal(err)
}
