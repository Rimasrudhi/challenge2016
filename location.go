package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

// stores locations
var locationMap = make(map[string]bool)

// read and store locations from CSV
func ReadCSV(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	// process each row, skipping the header
	for i, row := range records {
		if i == 0 || len(row) < 6 {
			continue
		}

		fullLocation := row[3] + "-" + row[4] + "-" + row[5] // City-State-Country
		locationMap[fullLocation] = true
	}

	fmt.Println("Total locations loaded:", len(locationMap))
}
