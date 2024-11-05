package log

import (
	"reflect"
	"testing"
)

// MockDriver is a mock implementation of the Driver interface for testing purposes.
type MockDriver struct {
	lastLogEntry LogEntry
}

// Log saves the log entry so we can inspect it later.
func (m *MockDriver) Log(entry LogEntry) error {
	m.lastLogEntry = entry
	return nil
}

// Test for NewLogger function
func TestNewLogger(t *testing.T) {
	mockDriver := &MockDriver{}
	logger := NewLogger(mockDriver)

	if logger.driver != mockDriver {
		t.Error("NewLogger() did not set the driver correctly")
	}
}

// Test for Logger.Log method
func TestLoggerLog(t *testing.T) {
	mockDriver := &MockDriver{}
	logger := NewLogger(mockDriver)

	level := INFO
	message := "Test message"
	attributes := map[string]interface{}{"key": "value"}

	// Call the Log method
	logger.Log(level, message, attributes)

	// Validate the created log entry
	if mockDriver.lastLogEntry.TransactionID == "" {
		t.Error("Expected non-empty TransactionID")
	}
	if mockDriver.lastLogEntry.Level != level {
		t.Errorf("Expected Level %v, got %v", level, mockDriver.lastLogEntry.Level)
	}
	if mockDriver.lastLogEntry.Message != message {
		t.Errorf("Expected Message %v, got %v", message, mockDriver.lastLogEntry.Message)
	}
	if !reflect.DeepEqual(mockDriver.lastLogEntry.Attributes, attributes) {
		t.Errorf("Expected Attributes %v, got %v", attributes, mockDriver.lastLogEntry.Attributes)
	}

	// Check if Timestamp is set (not zero)
	if mockDriver.lastLogEntry.Timestamp.IsZero() {
		t.Error("Timestamp is not set in the log entry")
	}
}
