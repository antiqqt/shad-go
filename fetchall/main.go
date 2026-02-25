//go:build !solution

package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	// urls := []string{
	// 	"https://example.com",
	// 	"https://httpbin.org/get",
	// 	"https://golang.org",
	// }

	urls := os.Args[1:]
	doneCh := make(chan error)
	now := time.Now()

	for _, u := range urls {
		go fetchURL(u, doneCh)
	}

	for range len(urls) {
		err := <-doneCh

		if err != nil {
			fmt.Printf("fetching error: %v", err)
		}
	}

	close(doneCh)

	fmt.Println("Elapsed:", time.Since(now))
}

func fetchURL(url string, doneCh chan<- error) {
	resp, err := http.Get(url)
	if err != nil {
		doneCh <- err
		return
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		doneCh <- err
		return
	}

	doneCh <- nil
}
