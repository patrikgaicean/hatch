package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

type ServiceInfo struct {
	Ip    string
	Dunno string
}

func main() {
	port := flag.Int("port", 8080, "Registry port")
	flag.Parse()

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx := context.Background()
	// err := client.Set(ctx, "foo", "bar", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }

	val, err := client.Get(ctx, "foo").Result()
	if err != nil {
		switch {
		case errors.Is(err, redis.Nil):
			fmt.Println("no value found")
		default:
			panic(err)
		}
	}
	fmt.Println("foo = ", val)

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		defer r.Body.Close()

		var dat2 ServiceInfo
		if err := json.Unmarshal([]byte(body), &dat2); err != nil {
			panic(err)
		}
		fmt.Printf("%+v", dat2)
	})

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", *port),
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Starting server on %s", srv.Addr)

	err = srv.ListenAndServe()
	log.Fatal(err)
}
