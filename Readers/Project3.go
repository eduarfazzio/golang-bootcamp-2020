package ReadCSV

import (
	//"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func ReadCSV() {
	// Open the file
	csvfile, err := os.Open("github/golang-bootcamp-2020/Files")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Question: %s Answer %s\n", record[0], record[1])
	}
}
