package api

import (
	"net/http"

	"github.com/dwarukira/findcare/internal/query"
	"github.com/dwarukira/findcare/pkg/txt"
	"github.com/gin-gonic/gin"
)

func GetErrors(router *gin.RouterGroup) {
	router.GET("/errors", func(c *gin.Context) {

		limit := txt.Int(c.Query("count"))
		offset := txt.Int(c.Query("offset"))

		if resp, err := query.Errors(limit, offset, c.Query("q")); err != nil {
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
