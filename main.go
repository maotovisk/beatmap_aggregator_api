package main

import (
	"beatmap_aggregator_api/config"
	"beatmap_aggregator_api/database"
	"log"
	"net/http"
	"os"
)

func main() {
	// Logging
	logFile, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(&LogWriter{
		StdOut:  os.Stdout,
		LogFile: logFile,
	})

	// Database
	database.InitDatabase()

	port := config.GetConfig().Web.Port

	if port == "" {
		port = "3000"
	}

	// HTTP Server
	r := Routes()

	log.Println("Server started on port:", port)
	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Printf("server error occured: %v", err)
	}
}
