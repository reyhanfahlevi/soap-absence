package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/reyhanfahlevi/soap-absence/server"
)

func main() {
	log.Printf("starting soap absence api")
	if err := server.Main(); err != nil {
		log.Fatal(err)
		return
	}
	log.Println("Program exited")
}
