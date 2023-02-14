package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://news.ycombinator.com/",
		"https://dr.dk",
		"https://golang.org",
		"https://old.reddit.com",
	}

	channel := make(chan string)

	for _, link := range links {
		go checkLink(link, channel)
	}

	// Waits for next response from channel
	for link := range channel {
		go func(currentLink string) {
			time.Sleep(time.Second)
			checkLink(currentLink, channel)
		}(link)
	}
}

func checkLink(link string, channel chan string) {
	fmt.Println("Getting", link, "...")
	_, error := http.Get(link)

	if error != nil {
		fmt.Println(link, "might be down")
		channel <- link
		return
	}

	fmt.Println(link, "is up")
	channel <- link
}
