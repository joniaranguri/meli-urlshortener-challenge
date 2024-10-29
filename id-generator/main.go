package main

import (
	"log"
	"os"
	"time"
)

const (
	batchSize = 60_000_000
	sleepTime = 1 * time.Hour
)

func main() {
	scope := os.Getenv("SCOPE")
	log.Printf("Running id-generator with scope %s", scope)
	db, err := NewUrlMappingDatabaseClient()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	for {
		currentId, err := getCurrentId(db)
		if err != nil {
			log.Printf("Error fetching current_id: %v", err)
			time.Sleep(sleepTime)
			continue
		}

		newCurrentId, err := generateAndInsertIds(db, currentId, batchSize)
		if err != nil {
			log.Printf("Error generating and inserting IDs: %v", err)
			time.Sleep(sleepTime)
			continue
		}

		if err := updateCurrentId(db, newCurrentId); err != nil {
			log.Printf("Error updating current_id: %v", err)
		}

		log.Printf("Batch completed. New current_id is %d", newCurrentId)

		if scope == "DEMO" {
			log.Println("Finishing run for Demo with 60 million registers")
			os.Exit(0)
		}
		time.Sleep(sleepTime)
	}
}
