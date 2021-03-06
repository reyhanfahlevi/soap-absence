package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/reyhanfahlevi/soap-absence/scheduler"
)

func main() {
	log.Printf("starting soap absence cron")
	if err := scheduler.Main(); err != nil {
		log.Fatal(err)
		return
	}
	log.Println("Program exited")
}
