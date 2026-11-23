package main

import (
	"encoding/csv"
	"os"
)

func loadRecipient(filepath string, ch chan Recipient) error {
	defer close(ch)

	f, err := os.Open(filepath)

	if err != nil {
		return err
	}

	r := csv.NewReader(f)
	defer f.Close()

	records, err := r.ReadAll()

	if err != nil {
		return err
	}

	for _, record := range records[1:] {
		// send -> consumers using channel
		ch <- Recipient{
			Name:  record[0],
			Email: record[1],
		}
	}
	return nil

}
