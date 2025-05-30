package crawler

// var visitedurls = make(map[string]bool)
// func Crawl(currenturl string, maxDepth int) {
// 	c := colly.NewCollector(
// 		colly.AllowedDomains("puang.in", "www.puang.in"),
// 		colly.MaxDepth(maxDepth),
// 	)
//
// 	c.OnHTML("title", func(e *colly.HTMLElement) {
// 		fmt.Println("Page Title:", e.Text)
// 	})
//
// 	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
// 		link := e.Request.AbsoluteURL(e.Attr("href"))
// 		// check if the current URL has already been visited
// 		if link != "" && !visitedurls[link] {
// 			visitedurls[link] = true
// 			fmt.Println("Found link:", link)
// 			e.Request.Visit(link)
// 		}
// 	})
//
// 	c.OnRequest(func(r *colly.Request) {
// 		fmt.Printf("Crawling %s\n", r.URL)
// 	})
//
// 	c.OnError(func(r *colly.Response, err error) {
// 		fmt.Printf("Error occurred while crawling %s: %v\n", r.Request.URL, err)
// 	})
//
// 	err := c.Visit(currenturl)
// 	if err != nil {
// 		fmt.Printf("Failed to visit %s: %v\n", currenturl, err)
// 	}
// }
