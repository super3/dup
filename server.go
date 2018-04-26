package main

import (
	"log"
	"fmt"
	"net/http"
	"time"
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/julienschmidt/httprouter"
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

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		seenThreshold := int64(10000)
		timeSince := fmt.Sprintf("%d", time.Now().Unix() - seenThreshold)
		rangeBy := redis.ZRangeBy{Min: timeSince, Max: fmt.Sprintf("%d", time.Now().Unix())}
		results, _ := client.ZRangeByScoreWithScores("nodes-active", rangeBy).Result();

		type node struct {
			NodeID interface{} `json:"nodeID"`
			LastSeen float64 `json:"lastSeen"`
		}

		nodes := []node{}

		for _, result := range results{
			nodes = append(nodes, node{NodeID: result.Member, LastSeen: result.Score})
		}

		resp := struct {
			Active []node `json:"active"`
		}{nodes}

		enc := json.NewEncoder(w)
		enc.SetIndent("", "    ")
		err := enc.Encode(resp)
		if err != nil {
			log.Println("error occured encoding json response", err)
		}
}

func Ping(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "hello, %s!\n", ps.ByName("address"))

		z := redis.Z{Score: float64(time.Now().Unix()), Member: ps.ByName("address")}
		client.ZAdd("nodes-active", z)
}

func main() {
	SetClient()

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/api/ping/:address", Ping)

	log.Fatal(http.ListenAndServe(":8080", router))
}
