# Telemetry Logging Package

This package provides a flexible and extensible logging solution for telemetry data, such as logs or transactions. Designed to be adaptable to different logging implementations (or "drivers"), it allows for seamless changes in logging backends without modifying the main application code. The library includes multiple log levels, supports various output formats, and is highly configurable.

## Features

- **Multiple Log Levels**: Supports `Debug`, `Info`, `Warning`, and `Error` levels for structured logging.
- **Flexible Output Drivers**: Output logs directly to the CLI, a simple text file, or a JSON file, with the ability to add new drivers easily.
- **Transaction-Styled Logs**: Each log entry can be associated with a `TransactionID` and additional attributes for enhanced tracking.
- **Extensible**: Supports adding new logging drivers without altering the core code.
- **Configurable**: Allows custom configurations through environment variables or a configuration file.
- **Ease of Use**: Runs with a sensible default configuration and minimal setup.

## Installation

Ensure you have Go installed and initialized in your project. Then, add the required packages by running:

```bash
go get github.com/caarlos0/env
go get github.com/joho/godotenv
```
## Usage

Environment Configuration
Create an .env file in your project directory to specify configuration values:

```bash
LOG_LEVEL=info
FILEPATH=./logs.json
DRIVER_TYPE=json
```
DRIVER_TYPE can be one of the following "cli", "json", "file". If not specified, the default is "cli".

## Setup
Initialize the logger service and start logging:

```bash
loggerService := InitLogger()
loggerService.Log(logger.DEBUG, "This is a debug message", map[string]interface{}{"CustomerId": "12345", "Environment": "production"})
```

## Directory Structure
```bash
.
├── .env                    # Environment configuration file
├── main.go                 # Application entry point
├── logger/
│   ├── logger.go           # Core logger and interface definitions
│   └── log_entry.go        # Log entry struct and level definitions
├── drivers/
│   ├── cli_driver.go       # CLI driver implementation
│   ├── file_driver.go      # Plain text file driver
│   └── json_file_driver.go # JSON file driver
└── logs/
    └── logs.json           # Example JSON log file (if using JSON driver)
```
## Extending with Custom Drivers
To add a new driver:

Implement the Driver interface in a new file in the drivers/ directory.
Add the driver initialization logic in main.go.
For example, a new database driver would look like:

```bash
type DatabaseDriver struct { /* fields here */ }

func (d DatabaseDriver) Log(entry logger.LogEntry) {
    // Database-specific implementation
}
```

## Testing
Unit tests can be created for each driver to ensure they conform to the Driver interface. Use Go's built-in testing package:

```bash
go test ./...
```

## License
This project is licensed under the MIT License. See the LICENSE file for details.

## Contributions
Contributions are welcome! Feel free to submit a pull request or open an issue to suggest new features, report bugs, or ask questions.

This package simplifies telemetry logging, offering a consistent and flexible logging experience that is easy to configure and extend. Enjoy logging!

```css
This `README.md` provides a thorough overview of setup, usage, customization, and contribution guidelines for the telemetry logging package. Let me know if you'd like any further details!
```