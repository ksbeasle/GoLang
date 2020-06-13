package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

/*
Exercise 1.11 --
Try a larger amount of websites
What happens if a website doesn't respond? ~~Throws err and continues?
*/

//go build main.go
// ./main https://golang.org http://gopl.io https://godoc.org https://www.upcdatabase.com/ guardian.co.uk newsvine.com amazon.co.jp furl.net tripadvisor.com youtube.com freewebs.com digg.com networkadvertising.org shop-pro.jp friendfeed.com sitemeter.com flickr.com redcross.org goodreads.com amazon.de de.vu pbs.org unblog.fr macromedia.com

func main() {

	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
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

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() //Don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
