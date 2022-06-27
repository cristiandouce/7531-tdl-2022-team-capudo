package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AttachTo(child *gin.Engine, parent *gin.RouterGroup) {
	routes := child.Routes()

	for _, route := range routes {
		parent.Handle(route.Method, route.Path, route.HandlerFunc)
	}
}

func ReplyWithInternalServerError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"error": err.Error()})
}

func ReplyWithBadRequesterror(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": err.Error()})
}

func ReplyWithNotFoundError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"error": err.Error()})

}