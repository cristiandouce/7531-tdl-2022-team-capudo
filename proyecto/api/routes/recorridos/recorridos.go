package recorridos

import (
	"capudo/api/routes"
	"capudo/dataBase"
	"capudo/dataBase/parser"
	"capudo/logger"
	"capudo/model"

	"errors"

	"net/http"
	"sync"

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
	//Recibe una id de destino por parametros y retorna los recorridos
	router.GET("bicicletero_destino/:id", func(ctx *gin.Context) {
		bicicleteroId := ctx.Param("id")
		if bicicleteroId == "" {
			routes.ReplyWithBadRequesterror(ctx, errors.New("id de bicicletero invalido"))
			return
		}
		recorridos, err := dataBase.GetRecorridos()
		if err != nil {
			routes.ReplyWithInternalServerError(ctx, err)
			return
		}
		filteredRecorrido := FilterRecorridos(*recorridos, func(recorrido model.Recorrido) bool {
			return recorrido.GetIdEstacionDestino() == bicicleteroId
		})
		ctx.JSON(http.StatusOK, filteredRecorrido)
	})
	//Recibe una id de origen por parametros y retorna los recorridos
	router.GET("bicicletero_origen/:id", func(ctx *gin.Context) {
		bicicleteroId := ctx.Param("id")
		if bicicleteroId == "" {
			routes.ReplyWithBadRequesterror(ctx, errors.New("id de bicicletero invalido"))
			return
		}
		recorridos, err := dataBase.GetRecorridos()
		if err != nil {
			routes.ReplyWithInternalServerError(ctx, err)
			return
		}

		filteredRecorrido := FilterRecorridos(*recorridos, func(recorrido model.Recorrido) bool {
			return recorrido.GetIdEstacionOrigen() == bicicleteroId
		})

		ctx.JSON(http.StatusOK, filteredRecorrido)
	})
	//Recibe una id de bicicletero por parametros y retorna los recorridos que tengan el id
	//como origen o destino
	//La consulta se realiza de forma concurrente
	router.GET("bicicletero_origen_destinoGO/:id", func(ctx *gin.Context) {
		wg := &sync.WaitGroup{}
		bicicleteroId := ctx.Param("id")
		if bicicleteroId == "" {
			routes.ReplyWithBadRequesterror(ctx, errors.New("id de bicicletero invalido"))
			return
		}
		recorridos, err := dataBase.GetRecorridos()
		if err != nil {
			routes.ReplyWithInternalServerError(ctx, err)
			return
		}
		var filteredRecorrido []model.Recorrido

		wg.Add(1)
		go func() {
			LoadFileteredRecorridos(&filteredRecorrido, *recorridos, func(recorrido model.Recorrido) bool {
				return recorrido.GetIdEstacionOrigen() == bicicleteroId
			})
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			LoadFileteredRecorridos(&filteredRecorrido, *recorridos, func(recorrido model.Recorrido) bool {
				return recorrido.GetIdEstacionDestino() == bicicleteroId
			})
			wg.Done()
		}()
		wg.Wait()
		ctx.JSON(http.StatusOK, filteredRecorrido)
	})
	//Recibe una id de bicicletero por parametros y retorna los recorridos que tengan el id
	//como origen o destino
	//La consulta NO se realiza de forma concurrente
	router.GET("bicicletero_origen_destino/:id", func(ctx *gin.Context) {
		bicicleteroId := ctx.Param("id")
		if bicicleteroId == "" {
			routes.ReplyWithBadRequesterror(ctx, errors.New("id de bicicletero invalido"))
			return
		}
		recorridos, err := dataBase.GetRecorridos()
		if err != nil {
			routes.ReplyWithInternalServerError(ctx, err)
			return
		}
		var filteredRecorrido []model.Recorrido

		LoadFileteredRecorridos(&filteredRecorrido, *recorridos, func(recorrido model.Recorrido) bool {
			return recorrido.GetIdEstacionOrigen() == bicicleteroId
		})

		LoadFileteredRecorridos(&filteredRecorrido, *recorridos, func(recorrido model.Recorrido) bool {
			return recorrido.GetIdEstacionDestino() == bicicleteroId
		})

		ctx.JSON(http.StatusOK, filteredRecorrido)
	})

	//Recibe una fecha por parametros y retorna los recorridos que se iniciaron posteriores a la misma
	//Ejemplo de consulta:
	// http://localhost:8000/api/recorridos/desde_fecha/2020-05-18/00:00:00
	router.GET("desde_fecha/:fecha/:hora", func(ctx *gin.Context) {
		fechaDesde := ctx.Param("fecha")
		horaDesde := ctx.Param("hora")
		if (fechaDesde == "" || horaDesde == "") {
			routes.ReplyWithBadRequesterror(ctx, errors.New("fecha invalida"))
			return
		}
		fechaDesdeTime, e := parser.ParseFechaHoraToTimeUTCError(fechaDesde, horaDesde, nil)		
		if(e != nil){
			logger.LogDebug("fecha invalida")
			routes.ReplyWithBadRequesterror(ctx, errors.New("fecha invalida"))
			return			
		}
		recorridos, err := dataBase.GetRecorridos()
		if err != nil {
			routes.ReplyWithInternalServerError(ctx, err)
			return
		}

		filteredRecorrido := FilterRecorridos(*recorridos, func(recorrido model.Recorrido) bool {
			return recorrido.GetFechaOrigenRecorrido().After(fechaDesdeTime) 
		})

		ctx.JSON(http.StatusOK, filteredRecorrido)
	})

}

func Attach(parent *gin.RouterGroup) {
	routes.AttachTo(submodule, parent)
}

func GetRoutes() gin.RoutesInfo {
	return submodule.Routes()
}
