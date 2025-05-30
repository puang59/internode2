package crawler

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/gocolly/colly"
)

func WebSearch(query string, maxResults int) ([]string, error) {
	c := colly.NewCollector()

	var result = []string{}

	c.OnHTML(".result__body", func(e *colly.HTMLElement) {
		link := e.ChildAttr(".result__a", "href")
		if link == "" {
			return
		}

		u, err := url.Parse(link)
		if err != nil {
			fmt.Printf("Error parsing URL %s: %v\n", link, err)
			return
		}

		finalUrl := link
		if u.Host == "duckduckgo.com" {
			uddg := u.Query().Get("uddg")
			if uddg != "" {
				var err error
				finalUrl, err = url.QueryUnescape(uddg)
				if err != nil {
					fmt.Printf("Error unescaping URL %s: %v\n", uddg, err)
					return
				}
			}
		}

		result = append(result, finalUrl)
		if len(result) >= maxResults {
			c.Visit("")
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("\nSearching for %s ...\n\n", query)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error occurred while searching for %s: %v\n", query, err)
	})

	startTime := time.Now()
	searchURL := fmt.Sprintf("https://html.duckduckgo.com/html/?q=%s", url.QueryEscape(query))
	err := c.Visit(searchURL)
	if err != nil {
		fmt.Printf("Failed to search for %s: %v\n", query, err)
		return []string{}, err
	}
	elapsed := time.Since(startTime)
	f, err := strconv.ParseFloat(fmt.Sprintf("%.2f", elapsed.Seconds()), 64)
	if err != nil {
		fmt.Printf("Error parsing elapsed time: %v\n", err)
		return []string{}, err
	}
	fmt.Printf("\nTime elapsed -> %.2f s\n", f)

	return result, nil
}
