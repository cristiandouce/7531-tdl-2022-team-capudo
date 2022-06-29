package parser

import (
	"capudo/model"
	"encoding/json"
	"io/ioutil"
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
		path (ruta del archivo .csv de bicicleteros)
	Returns:
		parBici (instancia de parser creada)
		e (el primer error en orden de aparición durante el parseo)
*/
func createParserBicicleterosPath(path string) (parBici *parserBicicleteros, e error) {
	parserAux, eParGen := ioutil.ReadFile(path)
	if eParGen == nil {
		parBici = new(parserBicicleteros)
		var features []model.Feature

		e = json.Unmarshal(parserAux, &features)
		if e != nil {
			return nil, e
		}
		for _, feat := range features {
			x := ParserFeatureToBicicletero(feat)
			parBici.bicicleteros = append(parBici.bicicleteros, x)

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
	lock_bicicleteros.Lock()
	defer lock_bicicleteros.Unlock()
	if instanceParserBicicleteros == nil {
		instanceParserBicicleteros, e = createParserBicicleteros()
		if e != nil {
			return nil, e
		}
	}
	bicicleteros = &instanceParserBicicleteros.bicicleteros

	return bicicleteros, e
}
