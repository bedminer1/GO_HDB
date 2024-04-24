package convert

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type HDBRecord struct {
	Month string `json:"month"`
	Town string `json:"town"`
	FlatType string `json:"flat_type"`
	Block string `json:"block"`
	LeaseStart string `json:"lease_start"`
	RemainingLease string `json:"remaining_lease"`
	Price int `json:"price"`
}

// convert csv lines to array of structs
func createRecordList(data [][]string, townFilter string, flatTypeFilter string, priceFilter int) []HDBRecord {
	var recordList []HDBRecord
	for i, line := range data {
		if i > 0 {
			var rec HDBRecord
			for j, field := range line {
				switch j {
				case 0:
					rec.Month = field
				case 1:
					rec.Town = field
				case 2: 
					rec.FlatType = field
				case 3:
					rec.Block = field
				case 8:
					rec.LeaseStart = field
				case 9:
					rec.RemainingLease = field
				case 10:
					rec.Price, _ = strconv.Atoi(field)
				}
			}

			// filters
			if townFilter != "" && rec.Town != townFilter {
				continue
			}
			if flatTypeFilter != "" && rec.FlatType != flatTypeFilter {
				continue
			}
			if priceFilter != 0 && priceFilter < rec.Price {
				continue
			}

			recordList = append(recordList, rec)
			
		}
	}
	return recordList
}

func CsvToArray(townFilter string, flatTypeFilter string, priceFilter string) []HDBRecord {
	// open file
	f, err := os.Open("convert/input/2017data.csv")
	if err != nil {
		fmt.Printf("Failed to open file: %v", err)
	}

	// close file at the end
	defer f.Close()


	// read csv file
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		fmt.Printf("Failed to read file: %v", err)
	}

	// fix filters
	priceNumFilter, _ := strconv.Atoi(priceFilter)
	fixedFlatTypeFilter := strings.Replace(flatTypeFilter, "+", " ", -1)
	fixedTownFilter := strings.Replace(townFilter, "+", " ", -1)
	fmt.Println(priceNumFilter, fixedTownFilter, fixedFlatTypeFilter)

	// convert to arr
	recordList := createRecordList(data, fixedTownFilter, fixedFlatTypeFilter, priceNumFilter)

	return recordList
}