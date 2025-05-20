package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	startTime := time.Now()
	for i := 0; i < 1000; i++ {
		resp, err := client.Get("https://github.com/")
		if err != nil {
			panic(err)
		}

		err = resp.Body.Close()
		if err != nil {
			panic(err)
		}
	}
	endTime := time.Now()

	fmt.Println("Time taken:", endTime.Sub(startTime).String())
}
