package main

import (
	"log"
	"fmt"
	"net/http"
	"github.com/go-redis/redis"
)

var client *redis.Client

func SetClient() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err) // Output: PONG <nil>
	fmt.Println(client); // Output: Redis<localhost:6379 db:0>
}

func main() {
	SetClient()

	// err := client.Set("key", 0, 0).Err()
	// if err != nil {
	// 	panic(err)
	// }
	val, err := client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
			fmt.Fprintln(w, "Hello world")
			err = client.Incr("key").Err()
			if err != nil {
				panic(err)
			}

			val, err = client.Get("key").Result()
			if err != nil {
				panic(err)
			}
			fmt.Println("key", val)
		})

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server", err)
	}
}
