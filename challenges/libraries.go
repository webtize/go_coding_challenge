package challenges

import (
	"time"
)

// LogEntry represents a parsed log line.
type LogEntry struct {
	Timestamp time.Time
	Level     string
	Message   string
}

// ParseLogEntries parses a slice of log strings.
// Each log string is in the format: "RFC3339_TIMESTAMP LEVEL: Message"
// Example: "2023-10-27T10:00:00Z INFO: Application started"
//
// If a log line is invalid (invalid timestamp or format), return an error.
func ParseLogEntries(logs []string) ([]LogEntry, error) {
	// TODO: Implement this function
	return nil, nil
}
