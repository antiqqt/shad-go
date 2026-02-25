//go:build !solution

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fileNames := os.Args[1:]
	hmap := make(map[string]int, 0)

	for _, path := range fileNames {
		file, err := os.Open(path)
		if err != nil {
			log.Fatalf("file open error: %v", err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			hmap[scanner.Text()]++
		}

		if err := scanner.Err(); err != nil {
			log.Fatalf("file read error: %v", err)
		}
	}

	for k, v := range hmap {
		if v < 2 {
			continue
		}

		fmt.Printf("%d\t%s\n", v, k)
	}
}
