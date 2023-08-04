package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

type ServiceInfo2 struct {
	ip    string
	dunno string
	Other string
	Prop  string
}

type ServiceInfo struct {
	Ip    string
	Dunno string
}

func register(addr string) {
	payload := &ServiceInfo{
		Ip:    addr,
		Dunno: "hi there",
	}
	jsonData, _ := json.Marshal(payload)
	fmt.Println(string(jsonData))

	res, err := http.Post(
		"http://localhost:8080",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("status: %d - body: %s\n", res.StatusCode, body)
}

func unregister() {}
func heartbeat()  {}

func main() {
	p := ServiceInfo2{
		ip:    "123",
		dunno: "456",
	}

	fmt.Printf("p = %+v\n", p)
	r := ServiceInfo2{
		ip:    "12344",
		dunno: "45664",
		Other: "12323",
		Prop:  "ksmggs",
	}
	fmt.Printf("r = %+v\n", r)

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
	// register(fmt.Sprint(listener.Addr()))

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
