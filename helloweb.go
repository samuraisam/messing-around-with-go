package main

import (
	"net/http"
	"fmt"
	"log"
)

func main() {
	resp, err := http.Head("http://fah-web.stanford.edu/daily_team_summary.txt.bz2")
	if err != nil {
		log.Fatal(err)
	}
	for key, value := range resp.Header {
		if key == "Last-Modified" {
			fmt.Println("Found last modified header!")
			fmt.Println("It was", value)
		}
	}
}