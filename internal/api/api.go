package api

import (
	"github.com/dwarukira/findcare/internal/event"
	"github.com/dwarukira/findcare/internal/i18n"
	"github.com/dwarukira/findcare/pkg/txt"
	"github.com/gin-gonic/gin"
)

var log = event.Log

func logError(prefix string, err error) {
	if err != nil {
		log.Errorf("%s: %s", prefix, err.Error())
	}
}

func Abort(c *gin.Context, code int, id i18n.Message, params ...interface{}) {
	resp := i18n.NewResponse(code, id, params...)

	log.Debugf("api: abort %s with code %d (%s)", c.FullPath(), code, resp.String())

	c.AbortWithStatusJSON(code, resp)
}

func Error(c *gin.Context, code int, err error, id i18n.Message, params ...interface{}) {
	resp := i18n.NewResponse(code, id, params...)

	if err != nil {
		resp.Details = err.Error()
		log.Errorf("api: error %s with code %d in %s (%s)", txt.Quote(err.Error()), code, c.FullPath(), resp.String())
	}

	c.AbortWithStatusJSON(code, resp)
}
