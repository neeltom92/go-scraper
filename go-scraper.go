package main

import (
	"fmt"
	"os"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"log"
	"time"
)

func main() {
	// Check if the URL argument is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run your_program.go <URL>")
		return
	}

	// Get the URL to scrape from the command line argument
	targetURL := os.Args[1]

	// Initialize a Chrome browser instance on port 4444
	service, err := selenium.NewChromeDriverService("/usr/local/bin/chromedriver", 4444)

	if err != nil {
		log.Fatal("Error:", err)
	}

	defer service.Stop()

	// Configure the browser options
	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{
		Args: []string{
			"--headless",
			"--disable-blink-features=AutomationControlled", // Disable automation detection
			"--user-agent=YOUR_USER_AGENT",                   // Set a user agent
		},
	})

	// Create a new remote client with the specified options
	driver, err := selenium.NewRemote(caps, "")

	if err != nil {
		log.Fatal("Error:", err)
	}

	// Visit the target page (URL from command line argument)
	err = driver.Get(targetURL)
	if err != nil {
		log.Fatal("Error:", err)
	}

	// Implement a delay to wait for potential JavaScript execution or CAPTCHA challenges
	time.Sleep(10 * time.Second)

	// Retrieve the page raw HTML as a string
	html, err := driver.PageSource()
	if err != nil {
		log.Fatal("Error:", err)
	}
	fmt.Println(html)
}

