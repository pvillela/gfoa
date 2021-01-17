package wgin

import (
	"github.com/gin-gonic/gin"
	"github.com/pvillela/gfoa/pkg/web"
)

type Router = *gin.Engine

func AddGetRoute(router Router, route string, handler gin.HandlerFunc) {
	router.GET(route, handler)
}

func AddPostRoute0(router Router, route string, handler web.PseudoPostHandler) {
	ginHandler := PostHanderMaker0(handler)
	router.POST(route, ginHandler)
}

func AddPostRoute(router Router, route string, handler gin.HandlerFunc) {
	router.POST(route, handler)
}
