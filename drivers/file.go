package drivers

import (
	"os"

	"github.com/leo-andrei/telemetry/log"
)

type FileDriver struct {
	FilePath string
}

func (d FileDriver) Log(entry log.LogEntry) error {
	f, err := os.OpenFile(d.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(entry.String() + "\n")
	return err
}
