package controllers

import (
	"fmt"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

func QuickSearch(query string, maxResults int) ([]any, error) {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"),
		colly.AllowURLRevisit(),
	)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
		r.Headers.Set("Accept-Language", "en-US,en;q=0.5")
		r.Headers.Set("Cache-Control", "no-cache")
		r.Headers.Set("Pragma", "no-cache")
	})

	c.SetRequestTimeout(10 * time.Second)
	// c.Limit(&colly.LimitRule{
	// 	DomainGlob:  "*",
	// 	RandomDelay: 3 * time.Second,
	// })

	var result = []any{}

	c.OnHTML(".result__body", func(e *colly.HTMLElement) {
		link := e.ChildAttr(".result__a", "href")
		if link == "" {
			return
		}

		title := e.ChildText(".result__a")
		if title == "" {
			return
		}

		u, err := url.Parse(link)
		if err != nil {
			return
		}

		finalUrl := link
		if u.Host == "duckduckgo.com" {
			uddg := u.Query().Get("uddg")
			if uddg != "" {
				finalUrl, err = url.QueryUnescape(uddg)
				if err != nil {
					return
				}
			}
		}

		result = append(result, []any{title, finalUrl})
		if len(result) >= maxResults {
			c.Visit("")
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Crawling %s\n", r.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error occurred while crawling %s: %v\n", r.Request.URL, err)
	})

	if maxResults > 10 && maxResults <= 20 {
		for i := range make([]int, 2) {
			var searchURL string
			if i == 0 {
				searchURL = fmt.Sprintf("https://html.duckduckgo.com/html/?q=%s", url.QueryEscape(query))
			} else {
				searchURL = fmt.Sprintf("https://html.duckduckgo.com/html/?q=%s&b=1", url.QueryEscape(query))
			}
			err := c.Visit(searchURL)
			if err != nil {
				return nil, fmt.Errorf("failed to visit search page on iteration %d: %v", i, err)
			}
		}
	} else {
		err := c.Visit(fmt.Sprintf("https://html.duckduckgo.com/html/?q=%s", url.QueryEscape(query)))
		if err != nil {
			return nil, fmt.Errorf("failed to visit search page: %v", err)
		}
	}

	return result, nil
}

func QuickSearchController(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(400, gin.H{
			"error": "Query parameter 'q' is required",
		})
		return
	}

	result, err := QuickSearch(query, 20)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to perform quick search",
		})
		return
	}

	if len(result) == 0 {
		c.JSON(404, gin.H{
			"message": "No results found for the query",
		})
		return
	}

	c.JSON(200, gin.H{
		"query":   query,
		"results": result,
	})

}
