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
	"sync"
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

/*
SetupLogger initializes zerolog to write to both console and a file.
It sets up logging to both the console and a file, allowing for both real-time monitoring and persistent record-keeping.

Parameters:

- logFilePath: The path to the log file where logs will be written.

- logLevel: The minimum level of logs to be written (e.g., DebugLevel, InfoLevel, ErrorLevel).

Returns:

- zerolog.Logger: The configured logger instance.

- error: An error if the log file cannot be opened or created.
*/
func SetupLogger(logFilePath string, logLevel zerolog.Level) (zerolog.Logger, error) {
	// Open or create the log file
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return zerolog.Logger{}, err
	}

	// Console writer with human-friendly formatting
	consoleWriter := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}

	// Set global log level
	zerolog.SetGlobalLevel(logLevel)

	// Combine both writers
	multi := zerolog.MultiLevelWriter(consoleWriter, file)

	// Set global time format
	zerolog.TimeFieldFormat = time.RFC3339

	// Create the logger
	logger := zerolog.New(multi).With().Timestamp().Logger()

	// Set as the global logger
	log.Logger = logger

	return logger, nil
}

/*
ScrapeDataXpath scrapes data from a webpage using the provided XPath and URL.
It initializes a new Colly collector, listens for elements matching the XPath,
and returns the first matched element or an error if the scraping fails.

Parameters:

- xpath: A string representing the XPath query to locate elements on the webpage.

- url: A string representing the URL of the webpage to scrape.

Returns:

- *colly.XMLElement: A pointer to the first matched XMLElement.

- error: An error if the scraping process encounters an issue.
*/
func ScrapeDataXpath(xpath string, url string) (*colly.XMLElement, error) {
	// Create a new collector
	c := colly.NewCollector()
	element := &colly.XMLElement{}
	c.OnXML(xpath, func(e *colly.XMLElement) {
		element = e
	})
	c.Visit(url)
	return element, nil
}

type currency struct {
	xpath string
	url   string
	name  string
} // access members by currency.xpath
// The function that will be executed
func main() {
	logger, err := SetupLogger("app.log", zerolog.InfoLevel)
	startTime := time.Now() // Record start time
	if err != nil {
		panic(err)
	}
	logger.Info().Msg(FgCyan + "Main function started" + Reset)
	// defer logger.Info().Msg(FgCyan + "Main function ended" + Reset)
	defer func() {
		duration := time.Since(startTime)
		logger.Info().Msgf(FgCyan+"Main function ended. Execution time: %s"+Reset, duration)
	}()

	xpath := "/html/body/main/div[1]/div[1]/div[1]/div/div[2]/div/h3[1]/span[2]/span[1]"
	// Create currency list
	currencies := []currency{}
	var currencyMap = map[string]float32{}
	USD := currency{
		name:  "USD",
		xpath: xpath,
		url:   "https://www.tgju.org/profile/price_dollar_rl",
	}
	EUR := currency{
		name:  "EUR",
		xpath: xpath,
		url:   "https://www.tgju.org/profile/price_eur",
	}
	AED := currency{
		name:  "AED",
		xpath: xpath,
		url:   "https://www.tgju.org/profile/price_aed",
	}
	GoldOunce := currency{
		name:  "GoldOunce",
		xpath: xpath,
		url:   "https://www.tgju.org/profile/ons",
	}
	currencies = append(currencies, USD)
	currencies = append(currencies, EUR)
	currencies = append(currencies, AED)
	currencies = append(currencies, GoldOunce)

	for _, value := range currencies {
		// Create a WaitGroup to wait for goroutines to finish
		var wg sync.WaitGroup // Waitgroup default value is 0
		// Add a count of 1 to the WaitGroup
		wg.Add(1) // So now we have 1 goroutine to wait for
		go func() {
			defer wg.Done() // Decrement the counter when the goroutine completes
			element, err := ScrapeDataXpath(value.xpath, value.url)
			if err != nil {
				log.Err(err)
			} else {
				// Remove , from element.text
				text := element.Text
				// Remove comma from the text
				textWithoutComma := ""
				for _, r := range text {
					if r != ',' {
						textWithoutComma += string(r)
					}
				}
				// Convert textWithoutComma to float32
				var floatValue float32
				fmt.Sscan(textWithoutComma, &floatValue)
				currencyMap[value.name] = floatValue
				log.Info().Msgf("Value of %s is %f", value.name, floatValue)
			} // Code to be executed in the goroutine
		}()

		// Wait for the goroutine to finish
		wg.Wait() // Will block until the WaitGroup counter is 0
	}
}
