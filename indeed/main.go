package main

import (
	"WebScrapping/indeed/scrapper"
	"os"
	"strings"

	"github.com/labstack/echo"
)

const fileName string = "./jobs.csv"

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrapy", handleScrape)
	e.Logger.Fatal(e.Start(":1325"))
}

// show the home.html file to client
func handleHome(c echo.Context) error {
	return c.File("./home.html")
}

// give client scrapping csv.file
func handleScrape(c echo.Context) error {
	// cilent에게 csv 파일 제공 후 서버에서 해당 파일 삭제
	defer os.Remove(fileName)
	query := c.FormValue("query")
	query = strings.ToLower(scrapper.CleanString(query))
	// start scrapping!
	scrapper.Scrape(query)
	return c.Attachment(fileName, query+"_jobs.csv")
}
