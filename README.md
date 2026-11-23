# Email Dispatcher (go-campaign-runner)

A concurrent email campaign dispatcher in Go using a producer-consumer pattern with goroutines and channels. It loads recipients from a CSV file, populates an email template, and dispatches them via SMTP.

## Architecture

- **Producer (`Producer.go`)**: Reads recipients from `emails.csv` and sends them into a channel.
- **Consumer (`consumer.go`)**: Concurrent worker pool that receives recipients from the channel and sends emails over SMTP.
- **Entry Point (`main.go`)**: Coordinates the pipeline using `sync.WaitGroup`.
- **Template (`email.tmpl`)**: Template for personalized email content.

## Prerequisites

- Go installed
- Local SMTP testing tool such as [Mailpit](https://github.com/axllent/mailpit) or [MailHog](https://github.com/mailhog/MailHog) listening on `localhost:1025`

## Running

1. Start your local SMTP server (e.g., Mailpit / MailHog on port 1025).
2. Run the application:
   ```bash
   go run .
   ```
