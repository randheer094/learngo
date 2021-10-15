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
		"http://golang.org",
		"http://amazon.com",
	}

	//test(links)
	// testRepeat(links)
	testRepeat1(links)
}

func test(links []string) {
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		c <- fmt.Sprint(link, " might be down!")
		return
	}
	c <- fmt.Sprint(link, " is up!")
}

func testRepeat(links []string) {
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}
	for {
		go checkLinkRepeat(<-c, c)
	}
}

func checkLinkRepeat(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
	} else {
		fmt.Println(link, " is up!")
	}
	c <- link
}

func testRepeat1(links []string) {
	c := make(chan string)
	for _, link := range links {
		go checkLink(link, c)
	}
	for l := range c {
		go checkLinkRepeat1(l, c)
	}
}

func checkLinkRepeat1(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
	} else {
		fmt.Println(link, " is up!")
	}
	time.Sleep(15 * time.Second)
	c <- link
}
