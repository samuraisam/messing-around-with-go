package main

import (
	"database/sql"
	"fmt"
	"log"
)

import (
	_ "github.com/lxn/go-pgsql"
)

func main() {
	db, err := sql.Open("postgres", "dbname=gotest user=sam password=kfloughvel")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var msg string

	err = db.QueryRow("SELECT $1 || ' ' || $2", "Hello", "SQL").Scan(&msg)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("poop master")
	fmt.Println(msg)
}
