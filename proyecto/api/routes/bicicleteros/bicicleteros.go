package bicicleteros

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var submodule *gin.Engine

func init() {
	submodule = gin.New()

	router := submodule.Group("/bicicleteros")

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}

func Attach(parent *gin.Engine) {
	routes := submodule.Routes()

	for _, r := range routes {
		parent.Handle(r.Method, r.Path, r.HandlerFunc)
	}
}
