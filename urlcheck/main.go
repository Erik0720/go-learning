package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
)

func main() {
	urls := os.Args[1:]
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go checkUrl(url, &wg)
	}

	wg.Wait()
}

func checkUrl(url string, wg *sync.WaitGroup) {
	resp, err := http.Get(url)
	defer wg.Done()
	if err != nil {
		fmt.Println("URL: " + url + ", Status: " + "ERROR")
		return
	}

	defer resp.Body.Close()
	status := resp.Status
	fmt.Println("URL: " + url + ", Status: " + status)

	return
}
