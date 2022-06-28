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
		// empezamos por parsear la query de la consulta
		query, err := ParseQuery(ctx)

		// de tener un error respondemos con 400 - Bad Request
		if err != nil {
			routes.ReplyWithBadRequesterror(ctx, err)
			return
		}

		// seguimos por obtener todos los registros de usuarios
		usuarios, err := dataBase.GetUsuarios()

		// de tener un error respondemos con 500 - Internal Server Error
		if err != nil {
			routes.ReplyWithInternalServerError(ctx, err)
			return
		}

		// Filtramos los resultados segun la query de consulta con nuestros filtros helper
		usuariosFiltrados := FiltrarPorFecha(usuarios, query.FechaAltaDesde, query.FechaAltaHasta)
		usuariosFiltrados = FiltrarPorEdad(&usuariosFiltrados, query.Edad, query.EdadDesde, query.EdadHasta)
		usuariosFiltrados = FiltrarPorGenero(&usuariosFiltrados, query.Genero)

		// en caso de no obtener ningun resultado, respondemos siempre con un array vacio
		if usuariosFiltrados == nil {
			usuariosFiltrados = make([]model.Usuario, 0)
		}

		// enviamos la respuesta!
		ctx.JSON(http.StatusOK, usuariosFiltrados)
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
