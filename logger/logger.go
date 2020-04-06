package logger

import (
	"cloud.google.com/go/logging"
	"context"
	"log"
	"os"
)

var (
	logName = "go-cloud-run"
)

func Info(text string) {
	write(text, logging.Info)
}

func Debug(text string) {
	write(text, logging.Debug)
}

func Error(err error) {
	write(err.Error(), logging.Error)
}

func write(text string, severity logging.Severity) {
	projectID := os.Getenv("PROJECT_ID")
	ctx := context.Background()
	if severity == 0 {
		severity = logging.Info
	}

	client, err := logging.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	logger := client.Logger(logName)
	logger.Log(logging.Entry{Payload: text, Severity: severity})
	if err := client.Close(); err != nil {
		log.Fatalf("Failed to close client: %v", err)
	}
}
