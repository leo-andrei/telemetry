package drivers

import (
	"fmt"
	"time"

	"github.com/leo-andrei/telemetry/log"
)

type CLIDriver struct{}

func (d CLIDriver) Log(entry log.LogEntry) error {
	fmt.Printf("[%s] [%s] %s: %s\n", entry.Timestamp.Format(time.RFC3339), entry.TransactionID, entry.Level, entry.Message)
	return nil
}
