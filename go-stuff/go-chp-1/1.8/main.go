package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

//Exercise 1.8 -- Modify the program to add 'http://' to each arg if it is missing (use strings.HasPrefix)
//go build main.go
//./main http://gopl.io
//./main gopl.io

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
			//fmt.Println(url)
		}
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
