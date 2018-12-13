package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func validityHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	series := r.URL.Query().Get("series")
	number := r.URL.Query().Get("number")

	if len(series) != 4 {
		errorHandlerBadRequest(w, "Length of series must be 4")
		return
	}
	if len(number) != 6 {
		errorHandlerBadRequest(w, "Length of number must be 6")
		return
	}
	if _, err := strconv.Atoi(series); err != nil {
		errorHandlerBadRequest(w, "Values of series and number should be numeric")
		return
	}
	if _, err := strconv.Atoi(number); err != nil {
		errorHandlerBadRequest(w, "Values of series and number should be numeric")
		return
	}

	if database.recordsNumber == 0 {
		errorHandlerInternalServerError(w, "Database is not ready")
		return
	}

	found, err := database.exists(series, number)
	if err != nil {
		errorHandlerInternalServerError(w, err.Error())
		return
	}

	writeValidityResult(w, !found)
}

func errorHandlerBadRequest(w http.ResponseWriter, errorMessage string) {
	w.WriteHeader(http.StatusBadRequest)
	_, _ = fmt.Fprintf(w, "{\"error\":\"%s\"}", errorMessage)
}

func errorHandlerInternalServerError(w http.ResponseWriter, errorMessage string) {
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = fmt.Fprintf(w, "{\"error\":\"%s\"}", errorMessage)
}

func writeValidityResult(w http.ResponseWriter, isValid bool) {
	validStr := "valid"
	if !isValid {
		validStr = "invalid"
	}
	_, _ = fmt.Fprintf(w, "{\"result\":\"%s\"}", validStr)
}
