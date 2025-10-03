/*
This is a template to build a go package. For more context and help watch this youtube video https://www.youtube.com/watch?v=sf7f4QGkwfE
Go packages can't be executed without a main package and a main function.
The use of packages are to share code between multiple projects or to organize code.

* Pakcage names,contastants,variables and funcitons that start with a capital letter are exported and can be used by other packages.

Generally a package is placed in its own directory with the same name as the package. So the cloger package would be placed in a directory named cloger and the file name will be cloger.go .

now to use this package in another package you would import it like this:

import "github.com/user/repository/path to cloger folder"
*/

package cloger

// Sample imports
import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)
// ANSI color codes for terminal output
const (
	Reset = "\033[0m"
	Bold  = "\033[1m"
	// Text Colors
	FgBlack   = "\033[30m"
	FgRed	 = "\033[31m"
	FgGreen   = "\033[32m"
	FgYellow  = "\033[33m"
	FgBlue	= "\033[34m"
	FgMagenta = "\033[35m"
	FgCyan	= "\033[36m"
	FgWhite   = "\033[37m"
	// Background Colors
	BgBlack   = "\033[40m"
	BgRed	 = "\033[41m"
	BgGreen   = "\033[42m"
	BgYellow  = "\033[43m"
	BgBlue	= "\033[44m"
	BgMagenta = "\033[45m"
	BgCyan	= "\033[46m"
	BgWhite   = "\033[47m"
)

/*
SetupLogger initializes zerolog to write to both console and a file.
*/
func SetupLogger(logFilePath string, logLevel zerolog.Level) (zerolog.Logger, error) {
	// Open or create the log file
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return zerolog.Logger{}, err
	}

	// Console writer with human-friendly formatting
	consoleWriter := zerolog.ConsoleWriter{
		Out:		os.Stdout,
		TimeFormat: time.RFC3339,
	}
	
	// Set global log level
	zerolog.SetGlobalLevel(logLevel)


	// Combine both writers
	multi := zerolog.MultiLevelWriter(consoleWriter, file)

	// Set global time format
	zerolog.TimeFieldFormat = time.RFC3339

	// Create the logger
	logger := zerolog.New(multi).With().Caller().Timestamp().Logger()

	// Set as the global logger
	log.Logger = logger

	return logger, nil
}
