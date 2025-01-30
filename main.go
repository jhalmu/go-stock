package main

import (
	"encoding/csv"
	"log"
	"os"
	_ "os/exec"
	"strconv"

	"github.com/a-h/templ"
	_ "github.com/a-h/templ"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
)

//go:generate go run github.com/a-h/templ/cmd/templ generate

type DividendData struct {
	Symbol   string  `json:"Symbol"`
	Year     int     `json:"Year"`
	Dividend float64 `json:"Dividend"`
}

func main() {
	app := fiber.New()

	app.Get("/", adaptor.HTTPHandler(templ.Handler(indexPage())))
	app.Get("/api/dividends", getDividends)

	log.Fatal(app.Listen(":3000"))
}

func getDividends(c *fiber.Ctx) error {
	data, err := readCSV("data.csv")
	if err != nil {
		return c.Status(500).SendString("Error reading CSV file")
	}
	return c.JSON(data)
}

func readCSV(filename string) ([]DividendData, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var data []DividendData
	for _, record := range records[1:] { // Skip header row
		year, _ := strconv.Atoi(record[1])
		dividend, _ := strconv.ParseFloat(record[2], 64)
		data = append(data, DividendData{
			Symbol:   record[0],
			Year:     year,
			Dividend: dividend,
		})
	}

	return data, nil
}
