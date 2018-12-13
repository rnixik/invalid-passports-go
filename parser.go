package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

var updatingInProcess bool

func parseSourceFile(invalidPassportsSourcePath string) int {
	log.Println("Parsing new records...")
	updatingInProcess = true

	csvFile, err := os.Open(invalidPassportsSourcePath)
	if err != nil {
		log.Print(err)
		updatingInProcess = false
		return 0
	}

	newInvalidPassports := make(map[uint16]map[uint32]bool)

	reader := csv.NewReader(bufio.NewReader(csvFile))
	i := 0
	recordsAdded := 0
	for {
		i += 1
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			updatingInProcess = false
			return 0
		}
		if len(line[0]) != 4 || len(line[1]) != 6 {
			continue
		}

		seriesInt, errSer := strconv.ParseInt(line[0], 10, 16)
		numberInt, errNum := strconv.ParseInt(line[1], 10, 32)

		if errSer != nil || errNum != nil {
			continue
		}

		seriesInt16 := uint16(seriesInt)
		numberInt32 := uint32(numberInt)

		if _, exists := newInvalidPassports[seriesInt16]; !exists {
			newInvalidPassports[seriesInt16] = make(map[uint32]bool)
		}

		newInvalidPassports[seriesInt16][numberInt32] = true
		recordsAdded += 1
	}

	if len(newInvalidPassports) > 0 {
		log.Printf("Parsed records: %d, added records: %d", i, recordsAdded)
		invalidPassports = newInvalidPassports
	} else {
		log.Println("No new records")
	}

	updatingInProcess = false

	return recordsAdded
}
