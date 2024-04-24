package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

type HDBRecord struct {
	Month string `json:"month"`
	Town string `json:"town"`
	FlatType string `json:"flat_type"`
	Block string `json:"block"`
	StoreyRange string `json:"storey_range"`
	FloorArea string `json:"floor_area"`
	Model string `json:"flat_model"`
	LeaseStart string `json:"lease_start"`
	RemainingLease string `json:"remaining_lease"`
	Price string `json:"price"`
}

func createRecordList(data [][]string) []HDBRecord {
	// convert csv to array of structs
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
				case 5:
					rec.StoreyRange = field
				case 6:
					rec.FloorArea = field
				case 7:
					rec.Model = field
				case 8:
					rec.LeaseStart = field
				case 9:
					rec.RemainingLease = field
				case 10:
					rec.Price = field
				}
			}
			recordList = append(recordList, rec)
		}
	}
	return recordList
}

func main() {
	// open file
	f, err := os.Open("input/2017data.csv")
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
	recordList := createRecordList(data)

	// convert to json
	jsonData, err := json.MarshalIndent(recordList, "", "  ")
	if err != nil {
		fmt.Printf("Failed to marshal into json: %v", err)
	}

	// write to json file
	os.WriteFile("output/2017data.json", jsonData, os.ModePerm)
}