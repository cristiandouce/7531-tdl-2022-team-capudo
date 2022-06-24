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

/*
	Returns:
		bicicleteros (arrego de bicicleteros)
		e (el primer error en orden de aparición durante el parseo)
	Nota: la función puede ser utilizada concurrentemente.
*/
func GetBicicleteros() (bicicleteros *[]model.Bicicletero, e error) {
	return parser.GetBicicleteros()
}