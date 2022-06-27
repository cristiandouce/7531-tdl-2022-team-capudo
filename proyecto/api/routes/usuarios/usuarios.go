package usuarios

import (
	"capudo/api/routes"
	"capudo/dataBase"
	"capudo/dataBase/parser"
	"capudo/model"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var submodule *gin.Engine

func init() {
	submodule = gin.New()

	router := submodule.Group("/usuarios")

	router.GET("", func(ctx *gin.Context) {
		usuarios, err := dataBase.GetUsuarios()

		// filteredUsers := dataBase.FilterUsuarios(*usuarios, func(u model.Usuario) bool {
		// 	return u.GetEdadUsuario() > 19
		// })

		if err != nil {
			routes.ReplyWithInternalServerError(ctx, err)
			return
		}

		ctx.JSON(http.StatusOK, usuarios)
	})

	router.GET("/:id", func(ctx *gin.Context) {
		userId, err := parser.ParseToUint32Error(ctx.Param("id"), nil)

		if err != nil {
			routes.ReplyWithBadRequesterror(ctx, err)
			return
		}

		usuarios, err := dataBase.GetUsuarios()

		if err != nil {
			routes.ReplyWithInternalServerError(ctx, err)
			return
		}

		var found bool
		var usuario model.Usuario

		for _, u := range *usuarios {
			if u.GetIdUsuario() == userId {
				found = true
				usuario = u
				break
			}
		}

		fmt.Println("USUARIO", usuario)

		if found {
			// Es necesario pasar la dirección de memoria
			// para que ctx.JSON invoke la interfaz de `MarshalJSON()`
			ctx.JSON(http.StatusOK, &usuario)
			return
		}

		routes.ReplyWithNotFoundError(ctx, errors.New("usuario no encontrado"))
	})
}

func Attach(parent *gin.RouterGroup) {
	routes.AttachTo(submodule, parent)
}

func GetRoutes() gin.RoutesInfo {
	return submodule.Routes()
}
