package main

import (
	"database/sql"
	"fmt"

	_ "github.com/glebarez/go-sqlite"
	"github.com/jhalmu/go-stock/database"
)

func main() {
	// connect to the SQLite database
	db, err := sql.Open("sqlite", "./my.db?_pragma=foreign_keys(1)")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	_ = db.QueryRow("select sqlite_version()")

	// create the countries table
	_, err = database.CreateTable(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	// read the CSV file
	countries, err := database.ReadCSV("./data/countries.csv")
	if err != nil {
		fmt.Println(err)
		return
	}

	// insert the data into the SQLite database
	for _, country := range countries {
		_, err := database.Insert(db, &country)
		if err != nil {
			fmt.Println(err)
			break
		}
	}

}
