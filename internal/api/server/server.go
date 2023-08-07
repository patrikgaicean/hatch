package server

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

func Serve(listener net.Listener) error {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello darkness my old friend \n"))
	})

	srv := &http.Server{
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Printf("Starting server on %s\n", listener.Addr())
	err := srv.Serve(listener)
	if err != nil {
		fmt.Println("here?")
		return err
	}

	return nil
}
