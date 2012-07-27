package main

import (
	// "net/http"
	"database/sql"
	"fmt"
	"log"
	"os"
	// "compress/bzip2"
)

// load the postgresql driver
import (
	_ "github.com/lxn/go-pgsql"
)

// get a database connection
func getDB() *sql.DB {
	conn, err := sql.Open("postgres", "dbname=gotest user=sam password=farty")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer conn.Close()

	return conn
}

// tests for whether or not the database for tracking downloads exists
func dbTableExists(name string) bool {
	db := getDB()

	var msg string
	q := "SELECT tablename FROM pg_tables WHERE schemaname='public' AND tablename=$1"
	err := db.QueryRow(q, name).Scan(&msg)

	return err != nil && len(msg) > 0
}

func createTable(tableName string, args string) {
	if !dbTableExists(tableName) {
		db := getDB()
		
	}
}

func main() {
	createTable("download-tracker", "lastDownload timestamp")
}
