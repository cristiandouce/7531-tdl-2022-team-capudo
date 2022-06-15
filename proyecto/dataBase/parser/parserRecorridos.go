package parser

import (
	"capudo/model"
	"sync"
)

var lock = &sync.Mutex{}

// Instancia del singleton del parser
var (
	instanceParserRecorridos *parserRecorridos
)

// Tipo de la instancia de recorrido
type parserRecorridos struct {
	recorridos []model.Recorrido
}

/*
	Params:
		row (slice con los del recorrido como strings)
	Returns:
		r (puntero a una instancia de Recorrido parseada a partir de row)
		e (el primer error de parseo en orden de aparición)
*/
func parseRowToRecorrido(row []string) (r *model.Recorrido, e error) {
	//id y duración
	id_recorrido := row[0]
	duracion_recorrido, e := ParseToUint32Error(row[1], e)
	//origen
	fecha_origen_recorrido, e := ParseTimeUTCError(row[2], e)
	id_estacion_origen := row[3]
	nombre_estacion_origen := row[4]
	direccion_estacion_origen := row[5]
	long_estacion_origen, e := ParseToFloat32Error(row[6], e)
	lat_estacion_origen, e := ParseToFloat32Error(row[7], e)
	//destino
	fecha_destino_recorrido, e := ParseTimeUTCError(row[8], e)
	id_estacion_destino := row[9]
	nombre_estacion_destino := row[10]
	direccion_estacion_destino := row[11]
	long_estacion_destino, e := ParseToFloat32Error(row[12], e)
	lat_estacion_destino, e := ParseToFloat32Error(row[13], e)
	//datos usuario y bicicleta
	id_usuario := row[14]
	modelo_bicicleta := row[15]
	if e == nil {
		r = model.CreateRecorrido(id_recorrido, duracion_recorrido,
			fecha_origen_recorrido, id_estacion_origen,
			nombre_estacion_origen, direccion_estacion_origen,
			long_estacion_origen, lat_estacion_origen,
			fecha_destino_recorrido, id_estacion_destino,
			nombre_estacion_destino, direccion_estacion_destino,
			long_estacion_destino, lat_estacion_destino,
			id_usuario, modelo_bicicleta)
	}
	return r, e
}

/*
	Params:
		path (ruta del archivo .csv de recorridos)
	Returns:
		parReco (instancia de parser creada)
		e (el primer error en orden de aparición durante el parseo)
*/
func createParserRecorridosPath(path string) (parReco *parserRecorridos, e error) {
	parserAux, eParGen := CreateParserGenerico(path)
	if eParGen == nil {
		parReco = new(parserRecorridos)
		for i := 1; i < len(parserAux.rows); i++ {
			rowParsed, eParRow := parseRowToRecorrido(parserAux.rows[i])
			if eParRow == nil && rowParsed != nil {
				parReco.recorridos = append(parReco.recorridos, *rowParsed)
			}
		}
	} else {
		e = eParGen
	}
	return parReco, e
}

/*
	Returns:
		parReco (instancia de parser creada a partir de la ruta por defecto)
		e (el primer error en orden de aparición durante el parseo)
*/
func createParserRecorridos() (parReco *parserRecorridos, e error) {
	return createParserRecorridosPath(PATH_RECORRIDOS)
}

/*
	Returns:
		recorridos (arrego de recorridos)
		e (el primer error en orden de aparición durante el parseo)
	Nota: la función utiliza una llave de exclusión internamente.
*/
func GetRecorridos() (recorridos *[]model.Recorrido, e error) {
	lock.Lock()
	defer lock.Unlock()
	if instanceParserRecorridos == nil {
		instanceParserRecorridos, e = createParserRecorridos()
		if e == nil {
			recorridos = &instanceParserRecorridos.recorridos
		}
	}
	return recorridos, e
}
