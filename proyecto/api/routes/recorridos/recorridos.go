package recorridos

import (
	"capudo/api/routes"
	"capudo/dataBase"
	"capudo/model"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var submodule *gin.Engine

func init() {
	submodule = gin.New()

	router := submodule.Group("/recorridos")

	router.GET("", func(ctx *gin.Context) {
		recorridos, err := dataBase.GetRecorridos()

		// filteredUsers := dataBase.Filterrecorridos(*recorridos, func(u model.Usuario) bool {
		// 	return u.GetEdadUsuario() > 19
		// })

		if err != nil {
			routes.ReplyWithInternalServerError(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, recorridos)
	})

	router.GET("/:id", func(ctx *gin.Context) {
		recorridoId := ctx.Param("id")

		if recorridoId == "" {
			routes.ReplyWithBadRequesterror(ctx, errors.New("id de recorrido invalido"))
			return
		}

		recorridos, err := dataBase.GetRecorridos()

		if err != nil {
			routes.ReplyWithInternalServerError(ctx, err)
			return
		}

		var found bool
		var recorrido model.Recorrido

		for _, r := range *recorridos {
			if r.GetIdRecorrido() == recorridoId {
				found = true
				recorrido = r
				break
			}
		}

		if found {
			// Es necesario pasar la dirección de memoria
			// para que ctx.JSON invoke la interfaz de `MarshalJSON()`
			ctx.JSON(http.StatusOK, &recorrido)
			return
		}

		routes.ReplyWithNotFoundError(ctx, errors.New("recorrido no encontrado"))
	})
}

func Attach(parent *gin.RouterGroup) {
	routes.AttachTo(submodule, parent)
}

func GetRoutes() gin.RoutesInfo {
	return submodule.Routes()
}
