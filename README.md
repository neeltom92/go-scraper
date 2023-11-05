# go-scraper

This is a Go program that uses the Selenium package to perform web scraping on a specified webpage. It can be used to extract the raw HTML content from a webpage.

[Why Selenium](https://www.browserstack.com/guide/web-scraping-using-selenium-python#:~:text=Web%20Scraping%20with%20Selenium%20allows%20you%20to%20gather%20all%20the%20required%20data%20using%20Selenium%20Webdriver%20Browser%20Automation.%20Selenium%20crawls%20the%20target%20URL%20webpage%20and%20gathers%20data%20at%20scale.%20This%20article%20demonstrates%20how%20to%20do%20web%20scraping%20using%20Selenium) ?

## Prerequisites

Before running this project, you need to have the following installed on your system:

1. Go: Make sure you have Go installed on your machine. You can download and install it from [https://golang.org/dl/](https://golang.org/dl/)

2. Ensure you have set the paths corretly [refer this blog on how to write Go code](https://go.dev/doc/code)

3. Chrome and ChromeDriver: You need Google Chrome and ChromeDriver installed on your system. You can download Chrome from [https://www.google.com/chrome/](https://www.google.com/chrome/) and ChromeDriver from [https://sites.google.com/a/chromium.org/chromedriver/downloads](https://sites.google.com/a/chromium.org/chromedriver/downloads).

4. Go Selenium Package: You'll need to install the Selenium package for Go. You can install it using the following command:

   ```shell
   go get github.com/tebeka/selenium

## How to Run Locally


  ```shell
1. Clone the repository:

git clone git@github.com:neeltom92/go-scraper.git
cd go-scraper

2. Open the main.go file and specify the correct path to your ChromeDriver executable in the code, Initialize a Chrome browser instance on port 4444, Replace "/path/to/chromedriver" with the actual path to your ChromeDriver executable.

service, err := selenium.NewChromeDriverService("/path/to/chromedriver", 4444)

3. Build the Go program:

go build

4. Run the program with the URL you want to scrape as a command-line argument. For example, to scrape "https://example.com," run:

./go-scraper https://google.com

```

## Example:


1. For example here we are scraping google.com:

<img src="https://github.com/neeltom92/go-scraper/assets/26869835/507be5cb-320b-45aa-9d20-1e3ec6bda4a5" width="800" alt="Image 1">
<img src="https://github.com/neeltom92/go-scraper/assets/26869835/5fe73e93-1c99-4260-9f59-296c15f7f4a7" width="600" alt="Image 2">
<img src="https://github.com/neeltom92/go-scraper/assets/26869835/a4588769-c93a-4924-8072-2cfeebc0ae2c" width="400" alt="Image 3">



