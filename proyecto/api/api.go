package api

import (
	"capudo/config"
	"net/http"

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

	// creamos la instancia root de nuestra Gin App (router)
	app = gin.Default()

	// eliminamos el `/` al final
	app.RedirectTrailingSlash = true

	// cargamos templates html
	app.LoadHTMLGlob("./**/templates/*")

	// definimos la ruta para el home
	app.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{})
	})

	// Cargamos nuestras API routes
	group := app.Group("/api")
	bicicleteros.Attach(group)
	recorridos.Attach(group)
	usuarios.Attach(group)
}

func Start() {
	app.Run(":" + config.Get("PORT"))
}
