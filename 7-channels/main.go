package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, url := range urls {
		go getStatus(url, c)
	}

	// Basic infinite loop
	// for {
	// 	go getStatus(<-c, c)
	// }

	// Loop catching c channel value changes
	for l := range c {
		go func(z string) {
			time.Sleep(5 * time.Second)
			getStatus(z, c)
		}(l)
	}
}

func getStatus(url string, c chan string) {
	_, err := http.Get(url)

	if err != nil {
		fmt.Println(url, "might be down!")
		c <- url
		return
	}

	fmt.Println(url, "is up!")
	c <- url
}
