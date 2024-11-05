package log

import (
	"reflect"
	"testing"
	"time"
)

// Test for LogLevel.String() method
func TestLogLevelString(t *testing.T) {
	tests := []struct {
		level    LogLevel
		expected string
	}{
		{DEBUG, "DEBUG"},
		{ERROR, "ERROR"},
		{WARNING, "WARNING"},
		{INFO, "INFO"},
	}

	for _, test := range tests {
		if result := test.level.String(); result != test.expected {
			t.Errorf("LogLevel.String() = %v, want %v", result, test.expected)
		}
	}
}

// Test for NewLogEntry function
func TestNewLogEntry(t *testing.T) {
	transactionID := "txn123"
	level := INFO
	message := "Test log entry"
	attributes := map[string]interface{}{"key1": "value1", "key2": 42}

	entry := NewLogEntry(transactionID, level, message, attributes)

	// Check if TransactionID, Level, Message, and Attributes are set correctly
	if entry.TransactionID != transactionID {
		t.Errorf("NewLogEntry().TransactionID = %v, want %v", entry.TransactionID, transactionID)
	}
	if entry.Level != level {
		t.Errorf("NewLogEntry().Level = %v, want %v", entry.Level, level)
	}
	if entry.Message != message {
		t.Errorf("NewLogEntry().Message = %v, want %v", entry.Message, message)
	}
	if !reflect.DeepEqual(entry.Attributes, attributes) {
		t.Errorf("NewLogEntry().Attributes = %v, want %v", entry.Attributes, attributes)
	}

	// Check if Timestamp is set (not zero)
	if entry.Timestamp.IsZero() {
		t.Error("NewLogEntry().Timestamp is not set")
	}
}

// Test for LogEntry.String() method
func TestLogEntryString(t *testing.T) {
	timestamp := time.Date(2023, 11, 1, 14, 0, 0, 0, time.UTC)
	entry := LogEntry{
		Timestamp:     timestamp,
		Level:         WARNING,
		Message:       "Warning message",
		TransactionID: "txn789",
		Attributes:    map[string]interface{}{"attribute": "value"},
	}

	expected := "2023-11-01T14:00:00Z [WARNING] [txn789] Warning message - map[attribute:value]"
	if result := entry.String(); result != expected {
		t.Errorf("LogEntry.String() = %v, want %v", result, expected)
	}
}
