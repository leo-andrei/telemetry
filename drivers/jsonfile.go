package drivers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/leo-andrei/telemetry/log"
)

type JSONFileDriver struct {
	FilePath string
}

// NewJSONFileDriver initializes a JSON file driver and opens the file for appending.
func NewJSONFileDriver(filePath string) JSONFileDriver {
	return JSONFileDriver{FilePath: filePath}
}

// Output implements the Driver interface for JSONFileDriver.
func (d JSONFileDriver) Log(entry log.LogEntry) error {
	// Read the existing contents of the file
	var entries []log.LogEntry
	fileData, err := ioutil.ReadFile(d.FilePath)

	// If the file has content, unmarshal it; otherwise, create a new array
	if err == nil && len(fileData) > 0 {
		if err := json.Unmarshal(fileData, &entries); err != nil {
			fmt.Println("Error unmarshalling existing file content:", err)
			return err
		}
	}

	// Append the new entry
	entries = append(entries, entry)

	// Marshal the updated array to JSON
	data, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return err
	}

	// Write the JSON array back to the file
	err = ioutil.WriteFile(d.FilePath, data, 0666)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	return nil
}
