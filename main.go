//Pay attention in order to run this script first
//you need to give this script a module name
//the format is <prefix>/<descriptive-text>
//we give this module a name with:
//$ go mod init <prefix>/<descriptive-text>
//after that we need to download the specified modules
//and list them in go.mod go dose that automaticlly by
//$ go mod tidy
//finnaly we can run the program by
//$ go run <gofile>
//To generate a binary file
//$ go build <gofile>

// Package declaration
package main

// import packages
import (
	"fmt"
	"os"
	"time"

	"github.com/gocolly/colly"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// ANSI color codes for terminal output
const (
	Reset = "\033[0m"
	Bold  = "\033[1m"
	// Text Colors
	FgBlack   = "\033[30m"
	FgRed     = "\033[31m"
	FgGreen   = "\033[32m"
	FgYellow  = "\033[33m"
	FgBlue    = "\033[34m"
	FgMagenta = "\033[35m"
	FgCyan    = "\033[36m"
	FgWhite   = "\033[37m"
	// Background Colors
	BgBlack   = "\033[40m"
	BgRed     = "\033[41m"
	BgGreen   = "\033[42m"
	BgYellow  = "\033[43m"
	BgBlue    = "\033[44m"
	BgMagenta = "\033[45m"
	BgCyan    = "\033[46m"
	BgWhite   = "\033[47m"
)

// SetupLogger initializes zerolog to write to both console and a file.
func SetupLogger(logFilePath string, level zerolog.Level) (zerolog.Logger, error) {
	//* Import these packages
	//* "os"
	//* "time"
	//* "github.com/rs/zerolog"
	//* "github.com/rs/zerolog/log"
	// Open or create the log file
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return zerolog.Logger{}, err
	}
	fileLogger := zerolog.New(file)
	// Console writer with human-friendly formatting
	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}

	// Combine both writers
	multi := zerolog.MultiLevelWriter(consoleWriter, fileLogger)

	// Set global time format
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// Set log level
	zerolog.SetGlobalLevel(level)

	// Create the logger
	logger := zerolog.New(multi).With().Timestamp().Caller().Logger()

	// Set as the global logger
	log.Logger = logger

	return logger, nil
}

/*
ScrapeDataXpath scrapes data from a webpage using the provided XPath and URL.
It initializes a new Colly collector, listens for elements matching the XPath,
and returns the first matched element or an error if the scraping fails.
*/
func ScrapeDataXpath(xpath string, url string) (*colly.XMLElement, error) {
	//* Make sure to import the necessary packages:
	//*	github.com/gocolly/colly
	//
	// Parameters:
	// - xpath: A string representing the XPath query to locate elements on the webpage.
	// - url: A string representing the URL of the webpage to scrape.
	//
	// Returns:
	// - *colly.XMLElement: A pointer to the first matched XMLElement.
	// - error: An error if the scraping process encounters an issue.

	// Create a new collector
	c := colly.NewCollector()
	element := &colly.XMLElement{}
	c.OnXML(xpath, func(e *colly.XMLElement) {
		element = e
	})
	c.Visit(url)
	return element, nil
}

// The function that will be executed
func main() {
	logger, err := SetupLogger("app.log", zerolog.InfoLevel)
	if err != nil {
		panic(err)
	}
	logger.Info().Msg(FgCyan + "Main function started" + Reset)
	defer logger.Info().Msg(FgCyan + "Main function ended" + Reset)
	element, err := ScrapeDataXpath("/html/body/main/div[1]/div[1]/div[1]/div/div[2]/div/h3[1]/span[2]/span[1]", "https://www.tgju.org/profile/price_dollar_rl")
	if err != nil {
		log.Error().Err(err)
	}
	fmt.Println(element.Text)
}
