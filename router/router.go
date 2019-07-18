package router

import (
	"net/http"
	"my-go-project/handler/sd"
	"my-go-project/router/middleware"

	"github.com/gin-gonic/gin"
)

// Load loads the middlewares, routes, handlers
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// middlewares
	g.Use(gin.Recovery())     // if error, recover api server
	g.Use(middleware.NoCache) // enforce browser do not use cache
	g.Use(middleware.Options) // cross domain
	g.Use(middleware.Secure)  // secure setting
	g.Use(mw...)
	// 404 Handler
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// The health check handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
