package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
		"http://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	for l := range c {
		go func() {
			time.Sleep(time.Second * 5)
			checkLink(l, c)
		}()
	}
}

func checkLink(l string, c chan string) {
	_, err := http.Get(l)

	if err != nil {
		fmt.Printf("%v Might be down\n", l)
		c <- l
	}

	fmt.Printf("%v is up\n", l)
	c <- l

}
