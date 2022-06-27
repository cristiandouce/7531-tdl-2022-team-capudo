package recorridos

import (
	"capudo/model"
)

type filterFuncRecorridos func(model.Recorrido) bool

func FilterRecorridos(recorridos []model.Recorrido, f filterFuncRecorridos) []model.Recorrido {
	var filtered []model.Recorrido
	for _, recorrido := range recorridos {
		if f(recorrido) {
			filtered = append(filtered, recorrido)
		}
	}
	return filtered
}

func LoadFileteredRecorridos(filtered *[]model.Recorrido, recorridos []model.Recorrido, f filterFuncRecorridos) {
	for _, recorrido := range recorridos {
		if f(recorrido) {
			*filtered = append(*filtered, recorrido)
		}
	}
}
