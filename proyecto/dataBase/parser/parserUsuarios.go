package parser

import (
	"capudo/model"
	"sync"
)

var lock_usuarios = &sync.Mutex{}

// Instancia del singleton del parser
var (
	instanceParserUsuarios *parserUsuarios
)

// Tipo de la instancia de bicicleteros
type parserUsuarios struct {
	usuarios []model.Usuario
}

/*
	Params:
		row (slice con los datos del usuario como strings)
	Returns:
		u (puntero a una instancia de Usuario parseada a partir de row)
		e (el primer error de parseo en orden de aparición)
*/
func parseRowToUsuario(row []string) (u *model.Usuario, e error) {
	//obtención de datos
	id_usuario, e := ParseToUint32Error(row[0], e)
	genero_usuario := row[1]
	edad_usuario, e := ParseToUint16Error(row[2], e)
	//Unifico la fecha
	fecha := row[3] + " " + row[4]
	fecha_alta, e := ParseTimeUserUTCError(fecha, e)

	if e == nil {
		u = model.CreateUsuario(id_usuario, genero_usuario, edad_usuario, fecha_alta)
	}
	return u, e
}

/*
	Params:
		path (ruta del archivo .csv de bicicleteros)
	Returns:
		parUser (instancia de parser creada)
		e (el primer error en orden de aparición durante el parseo)
*/
func createParserUsuariosPath(path string) (parUser *parserUsuarios, e error) {
	parserAux, eParGen := CreateParserGenerico(path)
	if eParGen == nil {
		parUser = new(parserUsuarios)
		for i := 1; i < len(parserAux.rows); i++ {
			rowParsed, eParRow := parseRowToUsuario(parserAux.rows[i])
			if eParRow == nil && rowParsed != nil {
				parUser.usuarios = append(parUser.usuarios, *rowParsed)
			}
		}
	} else {
		e = eParGen
	}
	return parUser, e
}

/*
	Returns:
		parUser (instancia de parser creada a partir de la ruta por defecto)
		e (el primer error en orden de aparición durante el parseo)
*/
func createParserUsuarios() (parUser *parserUsuarios, e error) {
	return createParserUsuariosPath(PATH_USUARIOS)
}

/*
	Returns:
		usuarios (arrego de bicicleteros)
		e (el primer error en orden de aparición durante el parseo)
	Nota: la función utiliza una llave de exclusión internamente.
*/
func GetUsuarios() (Usuarios *[]model.Usuario, e error) {
	if instanceParserUsuarios == nil {
		lock_usuarios.Lock()
		defer lock_usuarios.Unlock()
		instanceParserUsuarios, e = createParserUsuarios()
		if e == nil {
			Usuarios = &instanceParserUsuarios.usuarios
		}
	}
	return Usuarios, e
}
