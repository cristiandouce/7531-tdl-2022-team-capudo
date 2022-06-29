package bicicleteros

import (
	"capudo/api/routes"
	"capudo/database"
	"capudo/database/parser"
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
		// parseo query de la consulta
		query, err := ParseQuery(ctx)

		// si hay error respondemos con 400 - Bad Request
		if err != nil {
			routes.ReplyWithBadRequesterror(ctx, err)
			return
		}

		// obtención de todos los registros de recorridos sin filtrar
		bicicleteros, err := database.GetBicicleteros()

		// si hay error respondemos con 500 - Internal Server Error
		if err != nil {
			routes.ReplyWithInternalServerError(ctx, err)
			return
		}

		// Filtrado de resultados según la query de consulta con nuestros filtros helper
		bicicleterosFiltrados := FiltrarPorAnclajes(bicicleteros, query.Anclajes_t, query.Anclajes_max, query.Anclajes_min)

		if bicicleterosFiltrados == nil {
			bicicleterosFiltrados = make([]model.Bicicletero, 0)
		}
		ctx.JSON(http.StatusOK, bicicleterosFiltrados)
	})

	router.GET("/:id", func(ctx *gin.Context) {
		bicicleteroId, err := parser.ParseToUint32Error(ctx.Param("id"), nil)

		if err != nil {
			routes.ReplyWithBadRequesterror(ctx, err)
			return
		}

		bicicleteros, err := database.GetBicicleteros()

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
