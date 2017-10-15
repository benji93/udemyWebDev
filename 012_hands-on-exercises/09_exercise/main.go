package main


import (
	"encoding/csv"
	"log"
	"os"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	http.HandleFunc("/", csvHandler)
	http.ListenAndServe(":8080", nil)
}

func csvHandler(w http.ResponseWriter, r *http.Request) {
	records := parseCsv("table.csv")
	
	err := tpl.Execute(w, records)
	if err != nil {
		log.Fatalln(err)
	}
}

type Record struct {
	Date time.Time
	Open float64
}

func parseCsv(file string) []Record {
	csvFile, err := os.Open(file) // Grab the cvs file src
	if err != nil {
		log.Fatalln(err)
	}
	defer csvFile.Close() // Close the opening

	rdr	:= csv.NewReader(csvFile) // Read the csv file
	rows, err := rdr.ReadAll() 
	if err != nil {
		log.Fatalln(err)
	}

	var records []Record	
	for i, row := range rows { 
		if i == 0 { // We skip this line as it's the column titles
			continue 
		}
		date, _ := time.Parse("2006-01-02", row[0])
		open, _ := strconv.ParseFloat(row[1], 64)

		records = append(records, Record{
			Date: date,
			Open: open,
		})
	}
	return records
}