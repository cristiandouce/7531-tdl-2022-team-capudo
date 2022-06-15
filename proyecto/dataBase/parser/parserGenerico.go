package parser

import (
	"encoding/csv"
	"os"
)

// Es un slice de columnas parseadas como cadenas de caracteres
type parserGenerico struct {
	rows [][]string
}

/*
	Params:
		path (ruta del archivo .csv fuente)
	Returns:
		parGenerico (parser creado)
		e (el primer error en orden de aparición durante el parseo)
*/
func CreateParserGenerico(pathArchivo string) (p *parserGenerico, err error) {
	p = new(parserGenerico)
	archivo, err := os.Open(pathArchivo)
	if( err == nil){
		defer archivo.Close()
		csvReader := csv.NewReader(archivo)
		p.rows, err = csvReader.ReadAll()	
	} 
	return p, err
}

