package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

/*
Exercise 1.10 --
Find a website that produces a large amount of data. Investigate caching by running the program twice
To see whether the times were the same or how much its changed.
Is the content the same each time? ~~ maybe a second faster ~~
print the out to a file to be examined
*/

//go build main.go
// ./main https://golang.org http://gopl.io https://godoc.org https://www.upcdatabase.com/

func main() {

	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {

		go fetch(url, ch) //go routine
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) //receive channel
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) //send to channel ch
		return
	}

	//file to save output to -- I think thats what this exercise is asking to do
	f, err := os.Create("./output.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	nbytes, err := io.Copy(f, resp.Body)
	resp.Body.Close() //Don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
