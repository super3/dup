package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
		"time"
		"flag"
)
var privateKey string
var publicKey string

func ping() {
	response, err := http.Get(fmt.Sprintf("http://localhost:8080/api/ping/%s", publicKey))
	if err != nil {
			fmt.Printf("%s", err)
	} else {
			defer response.Body.Close()
			contents, err := ioutil.ReadAll(response.Body)
			if err != nil {
					fmt.Printf("%s", err)
			}
			fmt.Printf("%s\n", string(contents))
	}
}

func main() {
	flag.StringVar(&publicKey, "pubkey", "test", "description of the value in the flag")
	flag.StringVar(&privateKey, "privkey", "test", "description of the value in the flag")
	flag.Parse()

	timeout := 5 * time.Second
	for {
		ping()
		time.Sleep(timeout)
	}
}
