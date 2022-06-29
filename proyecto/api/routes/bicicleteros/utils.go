package bicicleteros

import (
	"capudo/model"
	"log"

	"github.com/gin-gonic/gin"
)

type QueryBicicleteros struct {
	Id         uint32  `form:"id" json:"id"`
	Codigo     string  `form:"codigo" json:"codigo"`
	Tipo       string  `form:"tipo" json:"tipo"`
	Ubicacion  string  `form:"ubicacion" json:"ubicacion"`
	Nombre     string  `form:"nombre" json:"nombre"`
	Latitud    float32 `form:"latitud" json:"latitud"`
	Longitud   float32 `form:"longitud" json:"longitud"`
	Anclajes_t uint16  `form:"anclajes_t" json:"anclajes_t"`
	// cuando no puede parsear la fecha, utilizamos `fecha.IsZero()` para saber si
	// fue inicializado o no antes de usarlo.
	Anclajes_max uint16 `form:"anclajes_max" json:"anclajes_max"`
	Anclajes_min uint16 `form:"anclajes_min" json:"anclajes_min"`
}

func ParseQuery(ctx *gin.Context) (query QueryBicicleteros, err error) {
	err = ctx.Bind(&query)
	return query, err
}

type filterFuncBicicleteros func(model.Bicicletero) bool

func FiltrarBicicleteros(bicicleteros *[]model.Bicicletero, f filterFuncBicicleteros) (filtered []model.Bicicletero) {
	for _, bicicletero := range *bicicleteros {
		if f(bicicletero) {
			filtered = append(filtered, bicicletero)
		}
	}
	return filtered
}

func FiltrarPorAnclajes(bicicleteros *[]model.Bicicletero, anclajes uint16, anclajes_max uint16, anclajes_min uint16) (filtrados []model.Bicicletero) {
	filtrados = FiltrarBicicleteros(bicicleteros, func(bicicletero model.Bicicletero) bool {
		if anclajes != 0 {
			return bicicletero.Anclajes_t == anclajes
		}

		if anclajes_min == 0 && anclajes_max == 0 {
			return true
		}

		if anclajes_min != 0 && anclajes_max != 0 {
			return bicicletero.Anclajes_t >= anclajes_min && bicicletero.Anclajes_t <= anclajes_max
		}

		if anclajes_min != 0 {
			return bicicletero.Anclajes_t >= anclajes_min
		}

		if anclajes_max != 0 {
			return bicicletero.Anclajes_t <= anclajes_max
		}

		log.Println("No matcheamos ningun filtro, ojo acá")
		return true
	})

	return filtrados
}
