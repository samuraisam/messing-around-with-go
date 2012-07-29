/*
 * Intent - in paralell, download a series of urls
 *  - timeout after 1 sec if the download has not completed
 *  - allow for failures, and move on (similar to timeout)
 *  - collect the stringified response from the fetched url
 */

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func main() {
	var urls = []string{
		"http://ssutch.org",
		"http://reddit.com",
		"http://news.ycombinator.com",
		"http://slashdot.org",
		"http://google.com",
		"http://samuraifilms.org",
		"http://gmail.com",
	}
	errors := make(chan string)
	messages := make(chan string)
	finished := make(chan bool)

	// downloaad all the things
	var waiter sync.WaitGroup
	go func() {
		// start downloading all urls
		for _, u := range urls {
			waiter.Add(1)                                // tell waiter we are downloading something
			go downloadUrl(u, messages, errors, &waiter) // fire off downloadUrl
		}
		// wait for everything to finish
		waiter.Wait()
		// when done, send finished
		finished <- true
	}()

	// loop forever waiting for responses
	done := false
	for !done {
		select {
		case err := <-errors:
			// we got an error, d'aw
			log.Print("error fetching: ", err)
		case msg := <-messages:
			// we got a message, a successful url, yay!
			log.Print("got message length: ", len(msg))
		case <-finished:
			// called when the waitgroup finishes (errors or not)
			log.Print("all finished!")
			done = true
		}
	}
}

func downloadUrl(url string, messages chan string, errors chan string, wg *sync.WaitGroup) {
	// TODO: timeout after 1 sec
	defer func() {
		wg.Done() // when we are done, tell the waitgroup that we are
	}()
	resp, err := http.Get(url) // get the url
	if err != nil {
		errors <- fmt.Sprintf("error getting: ", err) // send the url getting error
		return
	}
	val, err := ioutil.ReadAll(resp.Body) // read the body
	if err != nil {
		errors <- fmt.Sprintf("error reading: ", err) // send the read error
		return
	} else {
		messages <- string(val) // finally send the stringified response body
	}
}
