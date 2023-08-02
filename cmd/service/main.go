package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

func register() {
	res, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("status: %d - body: %s\n", res.StatusCode, body)
}

func unregister() {}
func heartbeat()  {}

func main() {
	// perform heartbeat
	go func() {
		for {
			time.Sleep(5 * time.Second)
			heartbeat()
		}
	}()

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal(err)
	}

	// register to service registry
	register()

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
	if err != nil {
		log.Fatal(err)
	}

	// on server <graceful> shutdown, de-register from service registry
	// TODO
}
