package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://golang.org",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	for i := 0; i < len(links) ; i++ {
		fmt.Println(<- c)
	}
}

func checkLink(link string, c chan string){
	body, err := http.Get(link)
	if(err != nil ){
		fmt.Println(link, body.Status)
		c <- body.Status
		return
	}
	fmt.Println(link, body.Status)
	c <- body.Status
}