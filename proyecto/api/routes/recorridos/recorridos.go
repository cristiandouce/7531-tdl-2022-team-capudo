package recorridos

import (
	"capudo/api/routes"
	"capudo/database"
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
		// parseo query de la consulta
		query, err := ParseQuery(ctx)

		// si hay error respondemos con 400 - Bad Request
		if err != nil {
			routes.ReplyWithBadRequesterror(ctx, err)
			return
		}

		// obtención de todos los registros de recorridos sin filtrar
		recorridos, err := database.GetRecorridos()

		// si hay error respondemos con 500 - Internal Server Error
		if err != nil {
			routes.ReplyWithInternalServerError(ctx, err)
			return
		}

		// Filtrado de resultados según la query de consulta con nuestros filtros helper
		recorridosFiltrados := FiltrarPorFecha(recorridos, query.FechaDesde, query.FechaHasta)
		recorridosFiltrados = FiltrarPorEstacionOrigen(&recorridosFiltrados, query.IDEstacionOrigen)
		recorridosFiltrados = FiltrarPorEstacionDestino(&recorridosFiltrados, query.IDEstacionDestino)
		recorridosFiltrados = FiltrarPorUsuario(&recorridosFiltrados, query.IDUsuario)

		if recorridosFiltrados == nil {
			recorridosFiltrados = make([]model.Recorrido, 0)
		}
		ctx.JSON(http.StatusOK, recorridosFiltrados)
	})

	router.GET("/:id", func(ctx *gin.Context) {
		recorridoId := ctx.Param("id")

		if recorridoId == "" {
			routes.ReplyWithBadRequesterror(ctx, errors.New("id de recorrido invalido"))
			return
		}

		recorridos, err := database.GetRecorridos()

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
