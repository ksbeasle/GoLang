package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

//Exercise 1.7 -- Use io.Copy() to copy the response body to os.stdout withouth requiring a large
// buffer to hold the entire stream, check io.Copy() error
//go build main.go
//./main http://gopl.io
func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		body, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", body)
	}
}
