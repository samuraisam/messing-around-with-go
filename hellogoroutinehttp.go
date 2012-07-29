package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	// list of urls to get
	u1 := "http://lightt.com"
	u2 := "http://google.com"
	u3 := "http://reddit.com"
	u4 := "http://samuraifilms.org"
	u5 := "http://gmail.com"
	u6 := "http://news.ycombinator.com"
	urls := []string{u1, u2, u3, u4, u5, u6}

	// make a channel 
	c := make(chan string)
	for _, u := range urls {
		go httpGet(u, c)
	}
	for i := 0; i < len(urls); i++ {
		<-c // wit for task to complete
		log.Print("A task completed: ", i)
	}
}

func httpGet(url string, c chan string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("error getting: ", err)
		c <- ""
		return
	}
	val, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("error reading: ", err)
		c <- ""
	} else {
		log.Print("got response: ", len(val), " bytes")
		c <- string(val)
	}
}
