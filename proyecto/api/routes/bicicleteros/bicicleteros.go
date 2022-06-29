package bicicleteros

import (
	"capudo/api/routes"
	"capudo/dataBase"
	"capudo/dataBase/parser"
	"capudo/model"
	"errors"

	"net/http"

	"github.com/gin-gonic/gin"
)

var submodule *gin.Engine

func init() {
	submodule = gin.New()

	router := submodule.Group("/bicicleteros")

	router.GET("", func(ctx *gin.Context) {
		bicicleteros, err := dataBase.GetBicicleteros()

		// filteredUsers := dataBase.Filterbicicleteros(*bicicleteros, func(u model.Usuario) bool {
		// 	return u.GetEdadUsuario() > 19
		// })

		if err != nil {
			routes.ReplyWithInternalServerError(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, bicicleteros)
	})

	router.GET("/:id", func(ctx *gin.Context) {
		bicicleteroId, err := parser.ParseToUint32Error(ctx.Param("id"), nil)

		if err != nil {
			routes.ReplyWithBadRequesterror(ctx, err)
			return
		}

		bicicleteros, err := dataBase.GetBicicleteros()

		if err != nil {
			routes.ReplyWithInternalServerError(ctx, err)
			return
		}

		var found bool
		var bicicletero model.Bicicletero

		for _, b := range *bicicleteros {
			if b.Id == bicicleteroId {
				found = true
				bicicletero = b
				break
			}
		}

		if found {
			// Es necesario pasar la dirección de memoria
			// para que ctx.JSON invoke la interfaz de `MarshalJSON()`
			ctx.JSON(http.StatusOK, &bicicletero)
			return
		}

		routes.ReplyWithNotFoundError(ctx, errors.New("bicicletero no encontrado"))
	})

}

func Attach(parent *gin.RouterGroup) {
	routes.AttachTo(submodule, parent)
}

func GetRoutes() gin.RoutesInfo {
	return submodule.Routes()
}
