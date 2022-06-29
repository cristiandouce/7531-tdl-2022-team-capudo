package recorridos

import (
	"capudo/model"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type filterFuncRecorridos func(model.Recorrido) bool

type QueryRecorridos struct {
	IDEstacionOrigen  string    `form:"id_estacion_origen" json:"id_estacion_origen"`
	IDEstacionDestino string    `form:"id_estacion_destino" json:"id_estacion_destino"`
	FechaDesde        time.Time `form:"fecha_desde" json:"fecha_desde" time_format:"2006-01-02"`
	FechaHasta        time.Time `form:"fecha_hasta" json:"fecha_hasta" time_format:"2006-01-02"`
	IDUsuario         string    `form:"id_usuario" json:"id_usuario"`
}

func ParseQuery(ctx *gin.Context) (query QueryRecorridos, err error) {
	err = ctx.Bind((&query))
	return query, err
}

func FiltrarRecorridos(recorridos *[]model.Recorrido, f filterFuncRecorridos) (filtered []model.Recorrido) {
	for _, recorrido := range *recorridos {
		if f(recorrido) {
			filtered = append(filtered, recorrido)
		}
	}
	return filtered
}

func FiltrarPorFecha(recorridos *[]model.Recorrido, fecha_desde time.Time, fecha_hasta time.Time) (filtered []model.Recorrido) {
	filtrados := FiltrarRecorridos(recorridos, func(recorrido model.Recorrido) bool {
		if fecha_desde.IsZero() && fecha_hasta.IsZero() {
			return true
		}
		if !fecha_desde.IsZero() && !fecha_hasta.IsZero() {
			return recorrido.GetFechaOrigenRecorrido().After(fecha_desde) && recorrido.GetFechaDestinoRecorrido().Before(fecha_hasta)
		}
		if !fecha_desde.IsZero() {
			return recorrido.GetFechaOrigenRecorrido().After(fecha_desde)
		}
		if !fecha_hasta.IsZero() {
			return recorrido.GetFechaDestinoRecorrido().Before(fecha_hasta)
		}
		log.Println("No matcheamos ningun filtro, ojo acá")
		return true
	})
	return filtrados
}

func FiltrarPorEstacionOrigen(recorridos *[]model.Recorrido, estacionOrigen string) (filtered []model.Recorrido) {
	filtrados := FiltrarRecorridos(recorridos, func(recorrido model.Recorrido) bool {
		if estacionOrigen != "" {
			return recorrido.IDEstacionOrigen == estacionOrigen
		}
		return true
	})
	return filtrados
}

func FiltrarPorEstacionDestino(recorridos *[]model.Recorrido, estacionDestino string) (filtered []model.Recorrido) {
	filtrados := FiltrarRecorridos(recorridos, func(recorrido model.Recorrido) bool {
		if estacionDestino != "" {
			return recorrido.IDEstacionDestino == estacionDestino
		}
		return true
	})
	return filtrados
}

func FiltrarPorUsuario(recorridos *[]model.Recorrido, id_usuario string) (filtered []model.Recorrido) {
	filtrados := FiltrarRecorridos(recorridos, func(recorrido model.Recorrido) bool {
		if id_usuario != "" {
			return recorrido.GetIdUsuario() == id_usuario
		}
		return true
	})
	return filtrados
}

/*
func FiltrarPorEstacionOrigenYDestino(recorridos *[]model.Recorrido, estacionOrigen string, estacionDestino string) (filtered []model.Recorrido){
	filtrados := FiltrarRecorridos(recorridos, func(recorrido model.Recorrido) bool{
		if((estacionOrigen != "") && (estacionDestino != "")){
			return (recorrido.IDEstacionOrigen == estacionOrigen) && (recorrido.IDEstacionDestino == estacionDestino)
		}
		return true
	})
	return filtrados
}*/
