package challenges

import (
	"errors"
	"time"
	"strings"
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
		var result []LogEntry

	for _, line := range logs {
		parts := strings.Fields(line)
		if len(parts) < 3 {
			return nil, errors.New("invalid log format")
		}

		time, err := time.Parse(time.RFC3339, parts[0])
		if err != nil {
			
			return nil, err
		}
		level := parts[1]
		if !strings.Contains(level, ":") {
			return nil, errors.New("missing ':'")
		}

		level = strings.Replace(level, ":", "", 1)
			
		result = append(result, LogEntry{
			Timestamp: time,
			Level: level,
		})
		
	}
	return result, nil
}
}
