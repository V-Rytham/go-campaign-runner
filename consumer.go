package main

import (
	"log"
	"net/smtp"
	"sync"
	"text/template"
	"time"
)

func emailWorker(id int, ch <-chan Recipient, wg *sync.WaitGroup, tmpl *template.Template, cfg Config, rateLimiter <-chan time.Time) {
	defer wg.Done()
	for recipient := range ch {
		if rateLimiter != nil {
			<-rateLimiter
		}

		addr := cfg.SMTPHost + ":" + cfg.SMTPPort
		// formattedMsg := fmt.Sprintf("To : %s\r\nSubject: Test Email\r\n\n%s\r\n", recipient.Email, "Just testing the email campaign")
		msg, err := executeTemplate(tmpl, recipient)

		if err != nil {
			log.Printf("[Worker %d] Template error for %s: %v\n", id, recipient.Email, err)
			continue
		}

		err = smtp.SendMail(addr, nil, cfg.SenderEmail, []string{recipient.Email}, []byte(msg))

		if err != nil {
			log.Printf("[Worker %d] Failed to send email to %s: %v\n", id, recipient.Email, err)
			continue
		}

		log.Printf("[Worker %d] Sent email to %s (%s)\n", id, recipient.Name, recipient.Email)
	}
}
