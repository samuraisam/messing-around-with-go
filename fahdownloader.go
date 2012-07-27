package main

import (
	// "net/http"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
	// "compress/bzip2"
)

// load the postgresql driver
import (
	_ "github.com/lxn/go-pgsql"
)

// get a database connection, must call own `defer {res}.Close()`
func getDB() *sql.DB {
	conn, err := sql.Open("postgres", "dbname=gotest user=sam password=farty")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	// defer conn.Close()

	return conn
}

// check for whether a database table exists
func dbTableExists(name string) bool {
	db := getDB()
	defer db.Close()

	// var msg string
	var nTables int64
	q := "SELECT COUNT(*) FROM pg_tables WHERE schemaname='public' AND tablename=$1"
	result, err := db.Query(q, name)
	result.Next()
	err = result.Err()

	if err != nil {
		log.Fatal("Error checking for table existence: ", err)
		os.Exit(1)
	}

	result.Scan(&nTables)

	return nTables > 0
}

// create a table if it does not exist
// `tableName` is the name of the table to check for existence
// `createStmt` is the actual CREATE TABLE... value
func createTable(tableName string, createStmt string) {
	if !dbTableExists(tableName) {
		fmt.Println("Creating table: ", createStmt)
		db := getDB()
		defer db.Close()
		_, err := db.Exec(createStmt)
		if err != nil {
			log.Fatal("Error creating table `", tableName, "` ", err)
			os.Exit(1)
		}
	}
}

// drop a table if it exists
func dropTable(tableName string) {
	if dbTableExists(tableName) {
		dropStmt := "DROP TABLE " + tableName
		fmt.Println("Dropping table: ", dropStmt)
		db := getDB()
		defer db.Close()
		_, err := db.Exec(dropStmt)
		if err != nil {
			log.Fatal("Error dropping table `", tableName, "` ", err)
			os.Exit(1)
		}
	}
}

var LINKS []string = {
	"http://fah-web.stanford.edu/daily_team_summary.txt.bz2"
	"http://fah-web.stanford.edu/daily_user_summary.txt.bz2"
}

func downloadStatsFiles() {
	db := getDB()
	defer db.Close()

	for _, val := range LINKS {
		q := `
			SELECT download_time 
			FROM downloads 
			WHERE link=$1 AND error=NULL ORDER BY download_time ASC
		`
		res, err := db.Query(q, val)
		if result.Next() {
			// there is at least one row, get the HEAD and make sure that we are not
			// going to download it more than required
			var lastModified time.Time
			err = result.Err()
			if err != nil {
				panic(fmt.Sprintf("query error: ", err))
			}
			res.Scan(&lastModified)
		} else {
			// there are no rows, just get it (first fetch)
		}
	}
}

func main() {
	createTable("downloads",
		`CREATE TABLE downloads (
			id SERIAL, 
			link text,
			download_time timestamp,
			last_modified timestamp with timezone,
			download_size_bytes int,
			download_duration float,
			error text
		)`)
	dropTable("downloads")
}
