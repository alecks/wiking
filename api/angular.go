package main

import (
	"path"
)

func init() {
	routes = append(routes, func() {
		e.File("/", path.Join("..", "angular", "dist", "wiking"))
		e.Static("/", path.Join("..", "angular", "dist", "wiking"))
	})
}
