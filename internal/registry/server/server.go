package server

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

func Serve(l net.Listener, r http.Handler) error {
	srv := &http.Server{
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Printf("Starting server on %s\n", l.Addr())
	err := srv.Serve(l)
	if err != nil {
		return err
	}

	return nil
}
