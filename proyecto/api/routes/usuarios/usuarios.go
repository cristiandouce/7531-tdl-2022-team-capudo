package usuarios

import (
	"capudo/api/routes"
	"capudo/dataBase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var submodule *gin.Engine

func init() {
	submodule = gin.New()

	router := submodule.Group("/usuarios")

	router.GET("", func(ctx *gin.Context) {
		usuarios, err := dataBase.GetUsuarios()

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error()})
			return
		}

		log.Println("USUARIOS", usuarios)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}

func Attach(parent *gin.RouterGroup) {
	routes.AttachTo(submodule, parent)
}

func GetRoutes() gin.RoutesInfo {
	return submodule.Routes()
}
