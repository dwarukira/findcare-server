package api

import (
	"net/http"

	"github.com/dwarukira/findcare/internal/query"
	"github.com/dwarukira/findcare/pkg/txt"
	"github.com/gin-gonic/gin"
)

// GET /api/v1/providers
func GetProviders(router *gin.RouterGroup) {
	router.GET("/providers", func(c *gin.Context) {

		limit := txt.Int(c.Query("count"))
		offset := txt.Int(c.Query("offset"))

		if resp, err := query.Providers(limit, offset); err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": txt.UcFirst(err.Error())})
			return
		} else {
			AddCountHeader(c, len(resp))
			AddLimitHeader(c, limit)
			AddOffsetHeader(c, offset)
			c.JSON(http.StatusOK, resp)
		}

	})
}
