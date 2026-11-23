package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"sync"
	"text/template"
	"time"
)

type Recipient struct {
	Name  string
	Email string
}
type Config struct {
	SMTPHost     string
	SMTPPort     string
	SenderEmail  string
	CSVPath      string
	TemplatePath string
	WorkerCount  int
	RateLimit    int
}

func main() {
	cfg := Config{}
	flag.StringVar(&cfg.SMTPHost, "smtp-host", "localhost", "SMTP host")
	flag.StringVar(&cfg.SMTPPort, "smtp-port", "1025", "SMTP port")
	flag.StringVar(&cfg.SenderEmail, "sender", "rythammaalpani44@gmail.com", "Sender email address")
	flag.StringVar(&cfg.CSVPath, "csv", "./emails.csv", "Path to recipients CSV")
	flag.StringVar(&cfg.TemplatePath, "tmpl", "email.tmpl", "Path to email template")
	flag.IntVar(&cfg.WorkerCount, "workers", 5, "Number of concurrent workers")
	flag.IntVar(&cfg.RateLimit, "rate", 10, "Emails per second limit (0 for unlimited)")
	flag.Parse()
	fmt.Println("Starting Email Dispatcher...")

	tmpl, err := template.ParseFiles(cfg.TemplatePath)
	if err != nil {
		log.Fatalf("failed to parse email template: %v", err)
	}

	var rateLimiter <-chan time.Time
	if cfg.RateLimit > 0 {
		interval := time.Second / time.Duration(cfg.RateLimit)
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		rateLimiter = ticker.C
	}

	var wg sync.WaitGroup
	recipientChan := make(chan Recipient, 50)

	go loadRecipient(cfg.CSVPath, recipientChan)

	for i := 1; i <= cfg.WorkerCount; i++ {
		wg.Add(1)
		go emailWorker(i, recipientChan, &wg, tmpl, cfg, rateLimiter)
	}
	wg.Wait()
	fmt.Println("All emails dispatched successfully!")
}

func executeTemplate(tmpl *template.Template, r Recipient) (string, error) {
	var tpl bytes.Buffer
	err := tmpl.Execute(&tpl, r)
	if err != nil {
		return "", err
	}
	return tpl.String(), nil
}
