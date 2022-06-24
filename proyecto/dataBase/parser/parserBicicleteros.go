package parser

import (
	"capudo/model"
	"sync"
)

var lock_bicicleteros = &sync.Mutex{}

// Instancia del singleton del parser
var (
	instanceParserBicicleteros *parserBicicleteros
)

// Tipo de la instancia de bicicleteros
type parserBicicleteros struct {
	bicicleteros []model.Bicicletero
}

/*
	Params:
		row (slice con los datos del bicicletero como strings)
	Returns:
		r (puntero a una instancia de Bicicletero parseada a partir de row)
		e (el primer error de parseo en orden de aparición)
*/
func parseRowToBicicletero(row []string) (b *model.Bicicletero, e error) {
	//obtención de datos
	lat, e := ParseToFloat32Error(row[0], e)
	long, e := ParseToFloat32Error(row[1], e)
	id, e := ParseToUint16Error(row[2], e)
	nombre_referencia := row[3]
	anio_de_ingreso, e := ParseToUint16Error(row[4], e)
	tipo := row[5]
	cantidad_bicicleteros, e := ParseToInt16WithEmptyError(row[6], e)
	ubicacion := row[7]
	clasificacion_lugar := row[8]
	calle := row[9]
	altura, e := ParseToUint16Error(row[10], e)
	calle_interseccion := row[11]
	barrio := row[12]
	comuna := row[13]
	codigo_postal := row[14]
	codigo_postal_argentino := row[15]
	if e == nil {
		b = model.CreateBicicletero(lat, long, id, nombre_referencia, 
									anio_de_ingreso, tipo, cantidad_bicicleteros,
									ubicacion, clasificacion_lugar,
									calle, altura, calle_interseccion, 
									barrio, comuna, codigo_postal,
									codigo_postal_argentino)
	}
	return b, e
}

/*
	Params:
		path (ruta del archivo .csv de bicicleteros)
	Returns:
		parBici (instancia de parser creada)
		e (el primer error en orden de aparición durante el parseo)
*/
func createParserBicicleterosPath(path string) (parBici *parserBicicleteros, e error) {
	parserAux, eParGen := CreateParserGenerico(path)
	if eParGen == nil {
		parBici = new(parserBicicleteros)
		for i := 1; i < len(parserAux.rows); i++ {
			rowParsed, eParRow := parseRowToBicicletero(parserAux.rows[i])
			if eParRow == nil && rowParsed != nil {
				parBici.bicicleteros = append(parBici.bicicleteros, *rowParsed)
			}
		}
	} else {
		e = eParGen
	}
	return parBici, e
}

/*
	Returns:
		parBici (instancia de parser creada a partir de la ruta por defecto)
		e (el primer error en orden de aparición durante el parseo)
*/
func createParserBicicleteros() (parBici *parserBicicleteros, e error) {
	return createParserBicicleterosPath(PATH_BICICLETEROS)
}

/*
	Returns:
		bicicleteros (arrego de bicicleteros)
		e (el primer error en orden de aparición durante el parseo)
	Nota: la función utiliza una llave de exclusión internamente.
*/
func GetBicicleteros() (bicicleteros *[]model.Bicicletero, e error) {
	if instanceParserBicicleteros == nil {
		lock_bicicleteros.Lock()
		defer lock_bicicleteros.Unlock()
		instanceParserBicicleteros, e = createParserBicicleteros()
		if e == nil {
			bicicleteros = &instanceParserBicicleteros.bicicleteros
		}
	}
	return bicicleteros, e
}