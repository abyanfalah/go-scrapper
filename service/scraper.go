package service

import (
	"fmt"
	"scrap/helper"
	"scrap/model"
	"time"

	"github.com/gocolly/colly/v2"
)

func Scrap() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.scrapingcourse.com"),
	)

	// called before an HTTP request is triggered
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	// triggered when the scraper encounters an error
	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong: ", err)
	})

	// fired when the server responds
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
	})

	// triggered when a CSS selector matches an element
	// c.OnHTML("a", func(e *colly.HTMLElement) {
	// 	// printing all URLs associated with the <a> tag on the page
	// 	fmt.Println(e.Attr("href"))
	// })

	c.OnHTML("a.next", CrawlNextPage)
	c.OnHTML("li.product", ScrapeProduct)

	// triggered once scraping is done (e.g., write the data to a CSV file)
	c.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, " scraped!")

		if err := helper.WriteProductsToCSV(model.ProductList, fmt.Sprintf("%d.csv", time.Now().Unix())); err != nil {
			fmt.Println("Error writing products to CSV:", err)
		}

	})

	c.Visit("https://www.scrapingcourse.com/ecommerce")
}
