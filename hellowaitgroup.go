/*
 * This package shows how to use sync.WaitGroup, to wait for a series of urls to download
 * in paralell in their own goroutines.
 * 
 * It is however probably an improper use of WaitGroup, as it seems one should not hope to
 * care much about the value that comes out of your goroutine.
 */
package main

import (
	"sync"
	"log"
	"net/http"
	"io/ioutil"
)

func main() {
	var urls = []string {
		"http://ssutch.org",
		"http://reddit.com",
		"http://news.ycombinator.com",
		"http://slashdot.org",
		"http://google.com",
		"http://samuraifilms.org",
		"http://gmail.com",
	}

	// used to synchronize the completion of the urls
	var wg sync.WaitGroup

	for _, u := range urls {
		// increment the WaitGroup counter
		wg.Add(1)
		// launch a goroutine to fetch the url
		go func(url string) {
			// get the url
			resp, err := http.Get(url)
			log.Print("getting: ", url)
			if err != nil {
				log.Fatal("error getting url: ", url, " error: ", err)
			}
			val, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal("error reading response: ", err)
			} else {
				// increment the number of successful requests to pull off the channel `c`
				log.Print("successful: ", len(val), " bytes")
			}
			log.Print("done: ", url)
			wg.Done()
		}(u)
	}

	// wait for the tasks to complete
	log.Print("waiting for tasks to complete")
	wg.Wait()
	log.Print("tasks are complete")
}