package main

import (
	"strconv"
)

func isInDatabase(invalidPassports *map[uint16]map[uint32]bool, series string, number string) (bool, error) {
	seriesInt, errSer := strconv.ParseInt(series, 10, 16)
	numberInt, errNum := strconv.ParseInt(number, 10, 32)

	if errSer != nil {
		return false, errSer
	}

	if errNum != nil {
		return false, errSer
	}

	seriesInt16 := uint16(seriesInt)
	numberInt32 := uint32(numberInt)

	found := false
	numberMap, seriesExists := (*invalidPassports)[seriesInt16]
	if seriesExists {
		_, found = numberMap[numberInt32]
	}

	return found, nil
}
