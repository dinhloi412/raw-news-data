package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

const URL = "https://vnexpress.net/"

func main() {
	c := colly.NewCollector()
	titles := make(map[string]bool)

	c.OnHTML("h3.title-news a", func(e *colly.HTMLElement) {
		title := e.Text
		if title != "" && !titles[title] {
			titles[title] = true
			if err := writeToFile(titles); err != nil {
				log.Printf("Failed to write titles to file: %v", err)
			}
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("-----------", r.URL.String())
	})

	err := c.Visit(URL)
	if err != nil {
		log.Fatal(err)
	}
}
