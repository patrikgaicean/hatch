package server

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

func Serve(listener net.Listener, router http.Handler) error {
	srv := &http.Server{
		Handler:      router,
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
