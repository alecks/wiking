package main

import (
	"path"

	"github.com/gin-gonic/gin"
)

func init() {
	routes = append(routes, route{
		path:   "/:fn",
		method: "GET",
		handlers: []gin.HandlerFunc{
			func(c *gin.Context) {
				fn := c.Param("fn")
				c.File(path.Join("..", "angular", "dist", "wiking", fn))
			},
		},
	})

	routes = append(routes, route{
		path:   "/",
		method: "GET",
		handlers: []gin.HandlerFunc{
			func(c *gin.Context) {
				c.Header("Content-Type", "text/html")
				c.File(path.Join("..", "angular", "dist", "wiking", "index.html"))
			},
		},
	})
}
