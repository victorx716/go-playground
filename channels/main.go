package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://coinmarketcap.com",
		"http://facebook.com",
		"http://golang.org",
		"http://amazon.com",
	}
	// free to pass around channel c
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "may be down")
		c <- "Might be down, possibly"
		return
	}
	fmt.Println(link, "is up!")
	c <- "It is actually up"
}
