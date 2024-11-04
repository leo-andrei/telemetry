package log

import (
	"fmt"
	"time"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	ERROR
	WARNING
	INFO
)

func (l LogLevel) String() string {
	return [...]string{"DEBUG", "ERROR", "WARNING", "INFO"}[l]
}

type LogEntry struct {
	Timestamp     time.Time
	Level         LogLevel
	Message       string
	TransactionID string
	Attributes    map[string]interface{}
}

func NewLogEntry(transactionID string, level LogLevel, message string, attributes map[string]interface{}) LogEntry {
	return LogEntry{
		Timestamp:     time.Now(),
		Level:         level,
		Message:       message,
		TransactionID: transactionID,
		Attributes:    attributes,
	}
}

func (entry LogEntry) String() string {
	return fmt.Sprintf("%s [%s] [%s] %s - %v", entry.Timestamp.Format(time.RFC3339), entry.Level, entry.TransactionID, entry.Message, entry.Attributes)
}
