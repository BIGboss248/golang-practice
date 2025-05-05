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

// Simple web scraping example using Colly and XPath
package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gocolly/colly"
)

func scrapeData(xpath string, url string) (*colly.XMLElement, error) {

	// scrapeData scrapes data from a webpage using the provided XPath and URL.
	// It initializes a new Colly collector, listens for elements matching the XPath,
	// and returns the first matched element or an error if the scraping fails.
	//
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

func main() {
	start := time.Now() // Record the start time

	var wg sync.WaitGroup // Waitgroup default value is 0
	// Add a count of 1 to the WaitGroup
	wg.Add(1) // So now we have 1 goroutine to wait for
	go func() {
		defer wg.Done() // Decrement the counter when the goroutine completes
		element, err := scrapeData("/html/body/main/div[1]/div[1]/div[1]/div/div[2]/div/h3[1]/span[2]/span[1]", "https://www.tgju.org/profile/geram24")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Element text:", element.Text)
	}()

	wg.Add(1) // Add another goroutine to wait format
	go func() {
		defer wg.Done() // Decrement the counter when the goroutine completes
		element, err := scrapeData("/html/body/main/div[1]/div[1]/div[1]/div/div[2]/div/h3[1]/span[2]/span[1]", "https://www.tgju.org/profile/price_dollar_rl")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Element text:", element.Text)
	}()

	// Wait for the goroutine to finish
	wg.Wait() // Will block until the WaitGroup counter is 0

	fmt.Println("\033[32mAll goroutines using wait group finished!\033[0m") // Print the message to the console

	elapsed := time.Since(start) // Calculate the elapsed time
	fmt.Printf("\033[33;4mExecution time: %s\033[0m\n", elapsed)
}
