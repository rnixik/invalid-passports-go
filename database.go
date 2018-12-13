package main

import "strconv"

type Database struct {
	invalidPassports       map[uint16]map[uint32]bool
	invalidPassportsBuffer map[uint16]map[uint32]bool
	recordsNumber          int64
	recordsNumberBuffer    int64
}

func NewDataBase() *Database {
	return &Database{
		make(map[uint16]map[uint32]bool, 1000000),
		make(map[uint16]map[uint32]bool, 1000000),
		0,
		0,
	}
}

func (db *Database) exists(series string, number string) (bool, error) {
	seriesInt, errSer := strconv.ParseInt(series, 10, 16)
	numberInt, errNum := strconv.ParseInt(number, 10, 32)

	if errSer != nil {
		return false, errSer
	}

	if errNum != nil {
		return false, errNum
	}

	seriesInt16 := uint16(seriesInt)
	numberInt32 := uint32(numberInt)

	found := false
	numberMap, seriesExists := db.invalidPassports[seriesInt16]
	if seriesExists {
		_, found = numberMap[numberInt32]
	}

	return found, nil
}

func (db *Database) addRecordToStoreBuffer(series string, number string) error {
	seriesInt, errSer := strconv.ParseInt(series, 10, 16)
	numberInt, errNum := strconv.ParseInt(number, 10, 32)

	if errSer != nil {
		return errSer
	}

	if errNum != nil {
		return errNum
	}

	seriesInt16 := uint16(seriesInt)
	numberInt32 := uint32(numberInt)

	if _, exists := db.invalidPassportsBuffer[seriesInt16]; !exists {
		db.invalidPassportsBuffer[seriesInt16] = make(map[uint32]bool)
	}

	db.invalidPassportsBuffer[seriesInt16][numberInt32] = true
	db.recordsNumberBuffer++

	return nil
}

func (db *Database) flushBufferToStore() {
	db.invalidPassports = db.invalidPassportsBuffer
	db.invalidPassportsBuffer = make(map[uint16]map[uint32]bool, 1000000)
	db.recordsNumber = db.recordsNumberBuffer
	db.recordsNumberBuffer = 0
}
