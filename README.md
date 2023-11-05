# go-scraper

This is a Go program that uses the Selenium package to perform web scraping on a specified webpage. It can be used to extract the raw HTML content from a webpage.

## Prerequisites

Before running this project, you need to have the following installed on your system:

1. Go: Make sure you have Go installed on your machine. You can download and install it from [https://golang.org/dl/](https://golang.org/dl/).

2. Chrome and ChromeDriver: You need Google Chrome and ChromeDriver installed on your system. You can download Chrome from [https://www.google.com/chrome/](https://www.google.com/chrome/) and ChromeDriver from [https://sites.google.com/a/chromium.org/chromedriver/downloads](https://sites.google.com/a/chromium.org/chromedriver/downloads).

3. Go Selenium Package: You'll need to install the Selenium package for Go. You can install it using the following command:

   ```shell
   go get github.com/tebeka/selenium

## How to Run Locally


  ```shell
1. Clone the repository:
  git clone git@github.com:neeltom92/go-scraper.git
  cd go-scraper

2. Open the main.go file and specify the correct path to your ChromeDriver executable in the code:

// Initialize a Chrome browser instance on port 4444, Replace "/path/to/chromedriver" with the actual path to your ChromeDriver executable.
  
   ```shell
   service, err := selenium.NewChromeDriverService("/path/to/chromedriver", 4444)

3. Build the Go program:

  ```shell
  go build


4. Run the program with the URL you want to scrape as a command-line argument. For example, to scrape "https://example.com," run:

  ```shell
  ./go-scraper https://google.com





