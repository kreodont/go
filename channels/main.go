package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string {"http://google.com", "http://facebook.com", "http://amazon.com", "http://golang.org",
	"http://fakesite.ooooo"}
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<- c)
	}
}

func checkLink(link string, channel chan string)  {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("There is an issue with ", link)
		channel <- "Might be down"
		return
	}
	fmt.Println(link, "OK")
	channel <- "Is up"
}