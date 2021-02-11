package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddContentTypeHeader adds a content type header to the response.
func AddContentTypeHeader(c *gin.Context, contentType string) {
	c.Header("Content-Type", contentType)
}

// AddCountHeader adds the actual result count to the response.
func AddCountHeader(c *gin.Context, count int) {
	c.Header("X-Count", strconv.Itoa(count))
}

// AddLimitHeader adds the max result count to the response.
func AddLimitHeader(c *gin.Context, limit int) {
	c.Header("X-Limit", strconv.Itoa(limit))
}

// AddOffsetHeader adds the result offset to the response.
func AddOffsetHeader(c *gin.Context, offset int) {
	c.Header("X-Offset", strconv.Itoa(offset))
}
