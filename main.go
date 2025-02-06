package main

import (
	"database/sql"
	"fmt"

	_ "github.com/glebarez/go-sqlite"
)

func main() {
	// connect to the SQLite database
	db, err := sql.Open("sqlite", "./my.db?_pragma=foreign_keys(1)")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// create the countries table
	_, err = CreateTable(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	// read the CSV file
	countries, err := ReadCSV("countries.csv")
	if err != nil {
		fmt.Println(err)
		return
	}

	// insert the data into the SQLite database
	for _, country := range countries {
		_, err := Insert(db, &country)
		if err != nil {
			fmt.Println(err)
			break
		}
	}

}
