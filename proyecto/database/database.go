package database

import (
	"capudo/database/parser"
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

/*
	Returns:
		bicicleteros (arrego de bicicleteros)
		e (el primer error en orden de aparición durante el parseo)
	Nota: la función puede ser utilizada concurrentemente.
*/
func GetUsuarios() (usuarios *[]model.Usuario, e error) {
	return parser.GetUsuarios()
}

// func SearchUsuarios(q QueryUsuario) []model.Usuario {
// 	var filtered []model.Usuario
// 	usuarios, err := GetUsuarios()

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for _, usuario := range *usuarios {
// 		resultAlta :=
// 	}

// 	return filtered
// }

// type filterFunc func(model.Usuario) bool

// type QueryUsuario struct {
// 	FromEdad uint16 `default: 1`
// 	ToEdad   uint16
// 	FromAlta time.Time
// 	ToAlta time.Time
// }

// func busquedaPorAlta (toAlta time.Time, fromAlta time.Time) {

// }
