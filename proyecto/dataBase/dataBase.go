package dataBase

import (
	"capudo/dataBase/parser"
	"capudo/model"
)

/*
	Returns:
		recorridos (arrego de recorridos)
		e (el primer error en orden de aparición durante el parseo)
	Nota: la función puede ser utilizada concurrentemente.
*/
func GetRecorridos() (recorridos *[]model.Recorrido, e error) {
	return parser.GetRecorridos()
}
