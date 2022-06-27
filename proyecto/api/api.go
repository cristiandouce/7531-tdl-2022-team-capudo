package api

import (
	"capudo/config"

	"capudo/api/routes/bicicleteros"
	"capudo/api/routes/recorridos"
	"capudo/api/routes/usuarios"

	"log"

	"github.com/gin-gonic/gin"
)

var app *gin.Engine

var logger = log.Default()

func init() {
	logger.SetPrefix("[CAPUDO-API] ")
	logger.Println("Iniciando el servidor API rest", config.Get("PORT"))

	app = gin.Default()
	app.RedirectTrailingSlash = true

	group := app.Group("/api")

	attach(group, bicicleteros.GetRoutes())
	recorridos.Attach(group)
	usuarios.Attach(group)
}

func attach(group *gin.RouterGroup, routes gin.RoutesInfo) {
	for _, r := range routes {
		group.Handle(r.Method, r.Path, r.HandlerFunc)
	}
}

func Start() {
	app.Run(":" + config.Get("PORT"))
}
