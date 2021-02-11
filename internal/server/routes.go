package server

import (
	"github.com/dwarukira/findcare/internal/api"
	"github.com/dwarukira/findcare/internal/config"
	"github.com/gin-gonic/gin"
)

func registerRoutes(router *gin.Engine, conf *config.Config) {
	// Enables automatic redirection if the current route can't be matched but a
	// handler for the path with (without) the trailing slash exists.
	router.RedirectTrailingSlash = true

	// Static assets like js, css and font files.
	router.Static("/static", conf.StaticPath())

	// JSON-REST API Version 1
	v1 := router.Group("/api/v1")
	{
		api.GetStatus(v1)
		api.GetErrors(v1)
		api.GetProviders(v1)
	}

}
