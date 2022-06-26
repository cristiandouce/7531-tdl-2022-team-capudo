package api

import (
	"capudo/config"

	"capudo/api/routes/bicicleteros"

	"fmt"

	"net/http"

	"github.com/gin-gonic/gin"
)

var server *gin.Engine

func init() {
	fmt.Println("Iniciando el servidor API rest", config.Get("PORT"))

	server = gin.Default()

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	bicicleteros.Attach(server)
}

func Start() {
	server.Run(":" + config.Get("PORT")) // listen and serve on 0.0.0.0:8080

}
