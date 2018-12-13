package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var database *Database
var invalidPassportsSourcePath string

func main() {
	var addr = flag.String("addr", ":8002", "http service address")
	var sourceFile = flag.String("source-file", "/tmp/list_of_expired_passports.csv", "path of csv with data")
	flag.Parse()

	invalidPassportsSourcePath = *sourceFile

	log.Printf("HTTP server using address %s", *addr)

	bindSignal()

	database = NewDataBase()

	err := parseSourceFile(database, invalidPassportsSourcePath)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Records in database: %d", database.recordsNumber)

	http.HandleFunc("/", validityHandler)

	err = http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func bindSignal() {
	signalsChannel := make(chan os.Signal, 1)
	signal.Notify(signalsChannel, syscall.SIGUSR1)

	go func() {
		for {
			sig := <-signalsChannel
			if sig == syscall.SIGUSR1 {
				log.Println("Received signal to update DB")

				err := parseSourceFile(database, invalidPassportsSourcePath)
				if err != nil {
					log.Println(err)
					continue
				}

				log.Printf("Records in database: %d", database.recordsNumber)
			}
		}
	}()
}
