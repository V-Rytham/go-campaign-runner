package main

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"
)

func email_worker(id int, ch <-chan Recipient, wg *sync.WaitGroup) {
	defer wg.Done()
	for recipient := range ch {
		smtpHost := "localHost"
		smtpPort := "1025"

		// formattedMsg := fmt.Sprintf("To : %s\r\nSubject: Test Email\r\n\n%s\r\n", recipient.Email, "Just testing the email campaign")
		msg, err := executeTemplate(recipient)

		if err != nil {
			fmt.Printf("Error")
			continue
		}

		err = smtp.SendMail(smtpHost+":"+smtpPort, nil, "rythammaalpani44@gmail.com", []string{recipient.Email}, []byte(msg))

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(id, " ", recipient.Email, " ", recipient.Name)
	}
}
