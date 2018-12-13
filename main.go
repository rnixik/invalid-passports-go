package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var invalidPassportsSourcePath string

var invalidPassports map[uint16]map[uint32]bool

func main() {
	var addr = flag.String("addr", ":8002", "http service address")
	var sourceFile = flag.String("source-file", "/tmp/list_of_expired_passports.csv", "path of csv with data")
	flag.Parse()

	invalidPassportsSourcePath = *sourceFile

	log.Printf("HTTP server using address %s", *addr)

	bindSignal()

	invalidPassports = make(map[uint16]map[uint32]bool)
	parseSourceFile(invalidPassportsSourcePath)

	http.HandleFunc("/", validityHandler)
	http.HandleFunc("/update-data", updateDataHandler)

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func bindSignal() {
	signalsChannel := make(chan os.Signal, 1)
	signal.Notify(signalsChannel, syscall.SIGUSR1)

	go func() {
		sig := <-signalsChannel
		if sig == syscall.SIGUSR1 {
			log.Println("Received signal to update DB")

			if updatingInProcess {
				log.Println("Already updating")
			} else {
				parseSourceFile(invalidPassportsSourcePath)
			}
		}
	}()
}
