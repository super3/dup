package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
		"time"
)

func ping() {
	response, err := http.Get("http://localhost:8080/api/ping/test")
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
	timeout := 5 * time.Second
	for {
		ping()
		time.Sleep(timeout)
	}
}
