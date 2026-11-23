# Email Dispatcher (go-campaign-runner)

A concurrent, rate-limited email campaign dispatcher in Go using a producer-consumer pattern with goroutines and channels. It streams recipients from a CSV file, populates an email template, and dispatches them via SMTP with configurable throughput and worker concurrency.

## Architecture

- **Producer (`producer.go`)**: Streams CSV recipient records row-by-row into a buffered channel with EOF detection and malformed row guarding.
- **Consumer (`consumer.go`)**: Concurrent worker pool that synchronizes against a shared rate-limiting ticker channel and dispatches emails over SMTP.
- **Entry Point (`main.go`)**: CLI flag parsing, template pre-parsing/caching, worker coordination, and `sync.WaitGroup` management.
- **Template (`email.tmpl`)**: Plain-text RFC 5322 MIME template with personalized headers (`From:`, `To:`, `Subject:`).

## Features

- **Pre-parsed Templates**: Template is parsed once at startup and executed concurrently across workers.
- **Memory Streaming**: Streams CSV row-by-row instead of reading the whole file into RAM.
- **Global Rate Limiting**: Shared `time.Ticker` channel throttles throughput across all concurrent workers.
- **Configurable CLI Flags**: Customise worker count, rate limit, sender, SMTP server, and input files.
- **Error Resilience**: Failed emails are logged with worker context and recipient details without aborting other workers.

## Prerequisites

- Go installed
- Local SMTP testing tool such as [Mailpit](https://github.com/axllent/mailpit) or [MailHog](https://github.com/mailhog/MailHog) listening on `localhost:1025`

## Usage

Run with default settings (5 workers, 10 emails/sec, Mailpit on `localhost:1025`):
```bash
go run .
```

### CLI Options

| Flag | Description | Default |
| :--- | :--- | :--- |
| `-rate` | Max emails dispatched per second (0 for unlimited) | `10` |
| `-workers` | Number of concurrent worker goroutines | `5` |
| `-sender` | Sender email address | `rythammaalpani44@gmail.com` |
| `-smtp-host` | SMTP server host | `localhost` |
| `-smtp-port` | SMTP server port | `1025` |
| `-csv` | Path to recipients CSV file | `./emails.csv` |
| `-tmpl` | Path to email template file | `email.tmpl` |

### Custom Example

```bash
go run . -workers=10 -rate=25 -csv=./emails.csv -smtp-port=1025
```
