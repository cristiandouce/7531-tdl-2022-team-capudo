package usuarios

import (
	"capudo/model"
	"log"
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
	FechaAltaDesde time.Time `form:"fecha_alta_desde" json:"fecha_alta_desde" time_format:"2006-01-02"`
	FechaAltaHasta time.Time `form:"fecha_alta_hasta" json:"fecha_alta_hasta" time_format:"2006-01-02"`
}

func ParseQuery(ctx *gin.Context) (query QueryUsuarios, err error) {
	err = ctx.Bind(&query)
	return query, err
}

type filterFuncUsuarios func(model.Usuario) bool

func FiltrarUsuarios(usuarios *[]model.Usuario, f filterFuncUsuarios) (filtered []model.Usuario) {
	for _, usuario := range *usuarios {
		if f(usuario) {
			filtered = append(filtered, usuario)
		}
	}
	return filtered
}

func FiltrarPorFecha(usuarios *[]model.Usuario, fecha_desde time.Time, fecha_hasta time.Time) (filtrados []model.Usuario) {
	filtrados = FiltrarUsuarios(usuarios, func(usuario model.Usuario) bool {
		if fecha_desde.IsZero() && fecha_hasta.IsZero() {
			return true
		}

		if !fecha_desde.IsZero() && !fecha_hasta.IsZero() {
			return usuario.GetFechaAlta().After(fecha_desde) && usuario.GetFechaAlta().Before(fecha_hasta)
		}

		if !fecha_desde.IsZero() {
			return usuario.GetFechaAlta().After(fecha_desde)
		}

		if !fecha_hasta.IsZero() {
			return usuario.GetFechaAlta().Before(fecha_hasta)
		}

		log.Println("No matcheamos ningun filtro, ojo acá")
		return true
	})

	return filtrados
}

func FiltrarPorEdad(usuarios *[]model.Usuario, edad uint16, edad_desde uint16, edad_hasta uint16) (filtrados []model.Usuario) {
	filtrados = FiltrarUsuarios(usuarios, func(usuario model.Usuario) bool {
		if edad != 0 {
			return usuario.GetEdadUsuario() == edad
		}

		if edad_desde == 0 && edad_hasta == 0 {
			return true
		}

		if edad_desde != 0 && edad_hasta != 0 {
			return usuario.GetEdadUsuario() >= edad_desde && usuario.GetEdadUsuario() <= edad_hasta
		}

		if edad_desde != 0 {
			return usuario.GetEdadUsuario() >= edad_desde
		}

		if edad_hasta != 0 {
			return usuario.GetEdadUsuario() <= edad_hasta
		}

		log.Println("No matcheamos ningun filtro, ojo acá")
		return true
	})

	return filtrados
}

func FiltrarPorGenero(usuarios *[]model.Usuario, genero string) (filtrados []model.Usuario) {
	filtrados = FiltrarUsuarios(usuarios, func(usuario model.Usuario) bool {
		if genero != "" {
			return usuario.GetGeneroUsuario() == genero
		}

		return true
	})

	return filtrados
}
