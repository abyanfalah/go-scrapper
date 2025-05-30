package service

import (
	"fmt"
	"scrap/model"
	"sync"

	"github.com/gocolly/colly/v2"
)

var visitedUrls sync.Map

func ScrapeProduct(h *colly.HTMLElement) {
	p := model.Product{}
	p.Url = h.ChildAttr("a", "href")
	p.Image = h.ChildAttr("img", "src")
	p.Name = h.ChildText(".product-name")
	p.Price = h.ChildText(".product-price")

	model.ProductList = append(model.ProductList, p)
}

func CrawlNextPage(h *colly.HTMLElement) {
	nextPageUrl := h.Attr("href")

	_, isVisited := visitedUrls.Load(nextPageUrl)
	if isVisited {
		fmt.Printf("%s has been visited before.", nextPageUrl)
		return
	}

	// mark url visited
	visitedUrls.Store(nextPageUrl, struct{}{})

	fmt.Printf("scraping %s ... \n", nextPageUrl)
	h.Request.Visit(nextPageUrl)
}
