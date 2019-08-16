package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

const maxCurrency = 10

func floodRequest() {
	// reused client object
	client := &http.Client{}
	endpoints := []string{"/", "/index", "/forbidden", "/badreq"}
	for {
		u := fmt.Sprintf("http://localhost:4433%s", endpoints[rand.Int()%4])
		req, _ := http.NewRequest(http.MethodGet, u, nil)
		if _, err := client.Do(req); err != nil {
			log.Printf("request error: %v", err)
		}
		time.Sleep(time.Millisecond * 250)
	}
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	for i := 0; i < maxCurrency; i++ {
		go floodRequest()
	}
	wg.Wait()
}
