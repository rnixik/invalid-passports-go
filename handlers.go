package main

import (
	"fmt"
	"log"
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

	if len(invalidPassports) < 1 {
		errorHandlerInternalServerError(w, "Database is not ready")
		return
	}

	found, err := isInDatabase(&invalidPassports, series, number)
	if err != nil {
		errorHandlerInternalServerError(w, err.Error())
		return
	}

	writeValidityResult(w, !found)
}

func updateDataHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if updatingInProcess {
		_, err := fmt.Fprintf(w, "{\"result\":\"already_in_process\"}")
		if err != nil {
			log.Println(err)
		}
		return
	}
	added := parseSourceFile(invalidPassportsSourcePath)
	_, err := fmt.Fprintf(w, "{\"result\":\"ok\",\"added\":%d}", added)
	if err != nil {
		log.Println(err)
	}
}

func errorHandlerBadRequest(w http.ResponseWriter, errorMessage string) {
	w.WriteHeader(http.StatusBadRequest)
	_, err := fmt.Fprintf(w, "{\"error\":\"%s\"}", errorMessage)
	if err != nil {
		log.Println(err)
	}
}

func errorHandlerInternalServerError(w http.ResponseWriter, errorMessage string) {
	w.WriteHeader(http.StatusInternalServerError)
	_, err := fmt.Fprintf(w, "{\"error\":\"%s\"}", errorMessage)
	if err != nil {
		log.Println(err)
	}
}

func writeValidityResult(w http.ResponseWriter, isValid bool) {
	validStr := "valid"
	if !isValid {
		validStr = "invalid"
	}
	_, err := fmt.Fprintf(w, "{\"result\":\"%s\"}", validStr)
	if err != nil {
		log.Println(err)
	}
}
