package crawler

import (
	"fmt"
	"github.com/gocolly/colly"
	"net/url"
)

var visitedurls = make(map[string]bool)
var result = []string{}

func isSameDomain(currenturl string, link string) bool {
	curParsed, err1 := url.Parse(currenturl)
	linkParsed, err2 := url.Parse(link)

	if err1 != nil || err2 != nil {
		return false
	}
	return curParsed.Host == linkParsed.Host && curParsed.Scheme == linkParsed.Scheme
}

func RecursiveCrawl(currenturl string, maxDepth int) ([]string, error) {
	c := colly.NewCollector(
		colly.MaxDepth(maxDepth),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))

		if link == "" || !isSameDomain(currenturl, link) {
			return
		}

		if link != "" && !visitedurls[link] {
			visitedurls[link] = true
			fmt.Println("Found link:", link)
			result = append(result, link)
			e.Request.Visit(link)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Crawling %s\n", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error occurred while crawling %s: %v\n", r.Request.URL, err)
	})

	err := c.Visit(currenturl)
	if err != nil {
		fmt.Printf("Failed to visit %s: %v\n", currenturl, err)
	}

	if len(visitedurls) == 0 {
		return nil, fmt.Errorf("no URLs found")
	}

	returnResult := result
	result = []string{}
	return returnResult, nil
}
