package usuarios

import (
	"time"

	"github.com/gin-gonic/gin"
)

type QueryUsuarios struct {
	Genero    string `form:"genero" json:"genero"`
	Edad      uint16 `form:"edad" json:"edad"`
	EdadDesde uint16 `form:"edad_desde" json:"edad_desde"`
	EdadHasta uint16 `form:"edad_hasta" json:"edad_hasta"`
	// cuando no puede parsear la fecha, utilizamos `fecha.IsZero()` para saber si
	// fue inicializado o no antes de usarlo.
	FechaAltaDesde time.Time `form:"fecha_alta_desde" json:"fecha_alta_desde"`
	FechaAltaHasta time.Time `form:"fecha_alta_hasta" json:"fecha_alta_hasta"`
}

func ParseQuery(ctx *gin.Context) (query QueryUsuarios, err error) {
	err = ctx.Bind(&query)
	return query, err
}
