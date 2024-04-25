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
	Model string `json:"model"`
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

func CsvToJSON(options FilterOptions) []interface{} {
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
	recordList, stats := createRecordList(data, fOptions)
	ret := []interface{} {recordList, stats}
	
	return ret
}

// convert csv lines to array of structs
func createRecordList(data [][]string, options fixedFilterOptions) ([]HDBRecord, map[string][]int) {
	var recordList []HDBRecord
	meanMap := make(map[string][]int) // eg "2017": [19200 (mean), 23 (count)]

	for i, line := range data {
		if i > 0 { // ignore header line
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
				case 7:
					rec.Model = field
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

			// updating mean
			year := rec.Month[:4]
			_, ok := meanMap[year]
			if !ok {
				meanMap[year] = []int{0, 0}
			}
			data := meanMap[year]
			mean, count := data[0], data[1]
			newMean := findMean(mean, count, rec.Price)
			meanMap[year] = []int{newMean, count + 1}
		}
	}

	// return more recent records first
	slices.Reverse(recordList)
	return recordList, meanMap
}

// findMean takes old mean, old count, new value and returns the mean
func findMean(oldMean int, oldCount int, newValue int) int {
	return (oldMean * oldCount + newValue) / (oldCount + 1)
}
