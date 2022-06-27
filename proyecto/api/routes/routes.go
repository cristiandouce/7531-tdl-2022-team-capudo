package routes

import "github.com/gin-gonic/gin"

func AttachTo(child *gin.Engine, parent *gin.RouterGroup) {
	routes := child.Routes()

	for _, route := range routes {
		parent.Handle(route.Method, route.Path, route.HandlerFunc)
	}
}
