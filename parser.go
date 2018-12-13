package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
)

var updatingInProcess bool

func parseSourceFile(database *Database, invalidPassportsSourcePath string) (err error) {
	if updatingInProcess {
		return errors.New("updating already is in process")
	}

	updatingInProcess = true
	log.Println("Parsing new records...")

	csvFile, err := os.Open(invalidPassportsSourcePath)
	if err != nil {
		updatingInProcess = false
		return
	}

	reader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			updatingInProcess = false
			return err
		}
		if len(line[0]) != 4 || len(line[1]) != 6 {
			continue
		}

		_ = database.addRecordToStoreBuffer(line[0], line[1])
	}

	database.flushBufferToStore()

	updatingInProcess = false

	return
}
