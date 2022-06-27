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

	bicicleteros.Attach(group)
	recorridos.Attach(group)
	usuarios.Attach(group)
}

func Start() {
	app.Run(":" + config.Get("PORT"))
}
