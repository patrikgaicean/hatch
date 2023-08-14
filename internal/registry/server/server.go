package server

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

func Serve(listener net.Listener) error {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("registryyy handler func"))
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
		return err
	}

	return nil
}
