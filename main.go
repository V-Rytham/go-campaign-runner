package main

import (
	"bytes"
	"fmt"
	"html/template"
	"sync"
)

type Recipient struct {
	Name  string
	Email string
}

func main() {
	var wg sync.WaitGroup
	RecipientChan := make(chan Recipient) //unbeffered channel

	go loadRecipient("./emails.csv", RecipientChan)

	workerCount := 5

	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go email_worker(i, RecipientChan, &wg)
	}
	wg.Wait()
	fmt.Println("Welcome to Email Dispatcher")
}

func executeTemplate(r Recipient) (string, error) {
	t, err := template.ParseFiles("email.tmpl")
	if err != nil {
		return "", err
	}
	var tpl bytes.Buffer

	error := t.Execute(&tpl, r)
	if error != nil {
		return "", error
	}
	return tpl.String(), nil
}
