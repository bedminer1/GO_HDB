package convert

import (
	"encoding/csv"
	"fmt"
	"os"
)

type HDBRecord struct {
	Month string `json:"month"`
	Town string `json:"town"`
	FlatType string `json:"flat_type"`
	Block string `json:"block"`
	LeaseStart string `json:"lease_start"`
	RemainingLease string `json:"remaining_lease"`
	Price string `json:"price"`
}

// convert csv lines to array of structs
func createRecordList(data [][]string, townFilter string) []HDBRecord {
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
					rec.Price = field
				}
			}
			// filter
			fmt.Println(townFilter)
			if rec.Town == townFilter {
				recordList = append(recordList, rec)
			}
		}
	}
	return recordList
}

func CsvToArray(townFilter string) []HDBRecord {
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

	// convert to arr
	recordList := createRecordList(data, townFilter)

	return recordList
}