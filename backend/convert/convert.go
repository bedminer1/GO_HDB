package convert

import (
	"encoding/csv"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type HDBRecord struct {
	Month string `json:"month"`
	Town string `json:"town"`
	FlatType string `json:"flatType"`
	LeaseStart string `json:"leaseStart"`
	RemainingLease string `json:"remainingLease"`
	Price int `json:"price"`
}

type FilterOptions struct {
	TownFilter string
	FlatTypeFilter string
	PriceFilter string
}

type fixedFilterOptions struct {
	TownFilter string
	FlatTypeFilter string
	PriceFilter int
}

func CsvToJSON(options FilterOptions) []HDBRecord {
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
	priceNumFilter, _ := strconv.Atoi(options.PriceFilter)
	
	fOptions := fixedFilterOptions {
		TownFilter: strings.Replace(options.TownFilter, "+", " ", -1),
		FlatTypeFilter: strings.Replace(options.FlatTypeFilter, "+", " ", -1),
		PriceFilter: priceNumFilter,
	}
	
	// convert to arr
	recordList := createRecordList(data, fOptions)
	
	return recordList
}

// convert csv lines to array of structs
func createRecordList(data [][]string, options fixedFilterOptions) []HDBRecord {
	var recordList []HDBRecord
	for i, line := range data {
		if i > 0 {
			var rec HDBRecord
			for j, field := range line {
				// match index in array(line) to field to be populated
				switch j {
				case 0:
					rec.Month = field
				case 1:
					rec.Town = field
				case 2: 
					rec.FlatType = field
				case 8:
					rec.LeaseStart = field
				case 9:
					rec.RemainingLease = field
				case 10:
					rec.Price, _ = strconv.Atoi(field)
				}
			}

			// filters
			if (options.TownFilter != "" && rec.Town != options.TownFilter) || 
			(options.FlatTypeFilter != "" && rec.FlatType != options.FlatTypeFilter) || 
			(options.PriceFilter != 0 && options.PriceFilter < rec.Price) {
				continue
			}

			recordList = append(recordList, rec)
		}
	}

	// return more recent records first
	slices.Reverse(recordList)
	return recordList
}