package database

import (
	"database/sql"
	"encoding/csv"
	"os"
	"strconv"
)

func CreateTable(db *sql.DB) (sql.Result, error) {
	sql := `CREATE TABLE IF NOT EXISTS countries (
        id INTEGER PRIMARY KEY,
        name     TEXT NOT NULL,
        population INTEGER NOT NULL,
        area INTEGER NOT NULL
    );`

	return db.Exec(sql)
}

type Country struct {
	Id         int
	Name       string
	Population int
	Area       int
}

func Insert(db *sql.DB, c *Country) (int64, error) {
	sql := `INSERT INTO countries (name, population, area) 
            VALUES (?, ?, ?);`
	result, err := db.Exec(sql, c.Name, c.Population, c.Area)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func Update(db *sql.DB, id int, population int) (int64, error) {
	sql := `UPDATE countries SET population = ? WHERE id = ?;`
	result, err := db.Exec(sql, population, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func Delete(db *sql.DB, id int) (int64, error) {
	sql := `DELETE FROM countries WHERE id = ?`
	result, err := db.Exec(sql, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func ReadCSV(filename string) ([]Country, error) {
	// Open the CSV file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read the CSV file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// Parse the CSV file
	var countries []Country
	for _, record := range records[1:] { // Skip header row
		population, err := strconv.Atoi(record[1])
		if err != nil {
			return nil, err
		}
		area, err := strconv.Atoi(record[2])
		if err != nil {
			return nil, err
		}
		country := Country{
			Name:       record[0],
			Population: population,
			Area:       area,
		}
		countries = append(countries, country)
	}

	return countries, nil
}
