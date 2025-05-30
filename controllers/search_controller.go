package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/puang59/wink/crawler"
)

func SearchController(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(400, gin.H{
			"error": "Query parameter 'q' is required",
		})
	}

	result, err := crawler.WebSearch(query, 10)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Failed to perform search",
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
