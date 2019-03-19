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
		"http://stackoverflow.com",
		"http://golang.com",
		"http://amazon.com",
	}
	// print := ""
	c := make(chan string)
	// print += <-c
	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(l string) {
			time.Sleep(time.Second * 5)
			checkLink(l, c)
		}(l)
		// fmt.Println(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	// str := link
	if err != nil {
		// fmt.Println(link, "might be down!")
		c <- "might be down!"
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
