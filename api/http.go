package main

import "github.com/gin-gonic/gin"

type route struct {
	path     string
	method   string
	handlers []gin.HandlerFunc
}

var r *gin.Engine

var routes []route

func serverListen() {
	gin.SetMode(gin.ReleaseMode)
	r = gin.Default()

	go setRoutes()

	r.Run()
}

// TODO: Improve this.
func setRoutes() {
	for _, v := range routes {
		if v.method == "GET" {
			r.GET(v.path, v.handlers...)
		} else if v.method == "POST" {
			r.POST(v.path, v.handlers...)
		} else if v.method == "OPTIONS" {
			r.OPTIONS(v.path, v.handlers...)
		} else if v.method == "DELETE" {
			r.DELETE(v.path, v.handlers...)
		} else if v.method == "PATCH" {
			r.PATCH(v.path, v.handlers...)
		} else if v.method == "PUT" {
			r.PUT(v.path, v.handlers...)
		} else if v.method == "HEAD" {
			r.HEAD(v.path, v.handlers...)
		}
	}
}
