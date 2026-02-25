//go:build !solution

package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	urls := os.Args[1:]

	for _, u := range urls {
		resp, err := http.Get(u)
		if err != nil {
			log.Fatalf("failed to fetch url: %v", u)
		}

		scanner := bufio.NewScanner(resp.Body)

		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			log.Fatalf("error reading response from %s: %v", u, err)
		}

		resp.Body.Close()
	}
}
