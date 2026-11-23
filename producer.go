package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func loadRecipient(filepath string, ch chan<- Recipient) error {
	defer close(ch)

	f, err := os.Open(filepath)
	if err != nil {
		log.Printf("Failed to open CSV file: %v\n", err)
		return err
	}
	defer f.Close()

	r := csv.NewReader(f)

	// Skip the header row (Name, Email)
	if _, err := r.Read(); err != nil {
		log.Printf("Failed to read CSV header: %v\n", err)
		return err
	}

	// Stream rows one by one
	for {
		record, err := r.Read()
		if err == io.EOF {
			break // Reached the end of file cleanly
		}
		if err != nil {
			log.Printf("Skipping malformed row due to error: %v\n", err)
			continue
		}

		// Guard against rows missing columns
		if len(record) < 2 {
			continue
		}

		ch <- Recipient{
			Name:  record[0],
			Email: record[1],
		}
	}

	return nil
}
