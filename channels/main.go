package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string {"http://google.com", "http://facebook.com", "http://amazon.com", "http://golang.org",
	"http://fakesite.ooooo"}
	for _, link := range links {
		go checkLink(link)
	}
}

func checkLink(link string)  {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("There is an issue with ", link)
		return
	}
	fmt.Println(link, "OK")
}