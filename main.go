package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"scrap/model"

	"github.com/gocolly/colly/v2"
)

func main() {
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

	c.OnHTML("li.product", func(h *colly.HTMLElement) {
		p := model.Product{}
		p.Url = h.ChildAttr("a", "href")
		p.Image = h.ChildAttr("img", "src")
		p.Name = h.ChildText(".product-name")
		p.Price = h.ChildText(".product-price")

		model.ProductList = append(model.ProductList, p)
	})

	// triggered once scraping is done (e.g., write the data to a CSV file)
	c.OnScraped(func(r *colly.Response) {
		fmt.Println(r.Request.URL, " scraped!")

		for _, v := range model.ProductList {
			v.PrintToScreen()
		}

		file, err := os.Create("products.csv")
		if err != nil {
			log.Fatalln("Failed to create output CSV file", err)
		}
		defer file.Close()

		// initialize a file writer
		writer := csv.NewWriter(file)

		// write the CSV headers
		headers := []string{
			"Url",
			"Image",
			"Name",
			"Price",
		}
		writer.Write(headers)

		// write each product as a CSV row
		for _, product := range model.ProductList {
			// convert a Product to an array of strings
			record := []string{
				product.Url,
				product.Image,
				product.Name,
				product.Price,
			}

			// add a CSV record to the output file
			writer.Write(record)
		}
		defer writer.Flush()

	})

	c.Visit("https://www.scrapingcourse.com/ecommerce")
}
