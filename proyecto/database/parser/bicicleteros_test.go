package parser

import (
	"capudo/model"
	"testing"
)

const PATH_TEST_BICICLETEROS string = "test_bicicleteros.csv"

func createDatosParserBicicleterosTest() (datos []model.Bicicletero) {
	datos = append(datos, *model.CreateBicicletero(-58.4569375558625, -34.6266557833432, 1,
		"LINEA A CARABOBO", 2011, "4U", 1, "VIA PUBLICA", "SUBTE", "RIVADAVIA AV.", 0,
		"CARABOBO AV.", "Flores", "Comuna 7", "", ""))
	datos = append(datos, *model.CreateBicicletero(-58.3804259230237, -34.611475939191, 107,
		"ASI", 2012, "4U", 1, "VIA PUBLICA", "EDIFICIOS PUBLICOS", "IRIGOYEN, BERNARDO DE", 272,
		"", "Monserrat", "Comuna 1", "1072", "C1072AAF"))
	datos = append(datos, *model.CreateBicicletero(-58.4263786839093, -34.5983935322703, 457,
		"ESC. PRIMARIA COMUN NÂ° 20 ROSARIO VERA PEÃ‘ALOZA", 2013, "4U", 1, "VIA PUBLICA", "ESCUELAS", "PRINGLES", 1165,
		"", "Almagro", "Comuna 5", "1183", "C1183AEU"))
	datos = append(datos, *model.CreateBicicletero(-58.4203948587377, -34.6547003849021, 472,
		"ESC. PRIMARIA COMUN NÂ° 02 DR. GENARO SISTO", 2013, "4U", 1, "VIA PUBLICA", "ESCUELAS", "TILCARA", 3365,
		"", "Nueva Pompeya", "Comuna 4", "1437", "C1437CYK"))
	datos = append(datos, *model.CreateBicicletero(-58.4868230292919, -34.5527978911952, 473,
		"ESC. DE EDUCACION MEDIA NÂ° 05/15Â°", 2013, "4U", 1, "INTERIOR", "ESCUELAS", "TRONADOR", 4134,
		"", "Saavedra", "Comuna 12", "1430", "C1430DMZ"))
	datos = append(datos, *model.CreateBicicletero(-58.4084262228404, -34.5865032086421, 478,
		"PLAZA LAS HERAS", 2013, "4U", 2, "VIA PUBLICA", "ESPACIOS VERDES", "DIAZ, CNEL. AV.", 0,
		"JUNCAL", "Recoleta", "Comuna 2", "", ""))
	datos = append(datos, *model.CreateBicicletero(-58.4660103877796, -34.5547388764391, 701,
		"Catering  Schuster", 2016, "1U", 2, "VIA PUBLICA", "Comercio amigo", "Amenabar", 3075,
		"", "NuÃ±ez", "Comuna 13", "1429", "C1429AEC"))
	datos = append(datos, *model.CreateBicicletero(-58.410203407127, -34.5920425550597, 702,
		"La MÃ¡quina de Jugar", 2016, "1U", 2, "VIA PUBLICA", "Comercio amigo", "SÃ¡nchez de Bustamante", 1645,
		"", "Recoleta", "Comuna 2", "", ""))
	datos = append(datos, *model.CreateBicicletero(-58.3738466790501, -34.6026955573037, 707,
		"Ministerio de Ambiente de NaciÃ³n", 2016, "1U", 2, "VIA PUBLICA", "EDIFICIOS PÃšBLICOS", "San MartÃ­n", 451,
		"", "San Nicolas", "Comuna 1", "", ""))
	datos = append(datos, *model.CreateBicicletero(-58.3989398022133, -34.6413855694075, 708,
		"Registro Nacional de las Personas RENAPER", 2016, "1U", 3, "VIA PUBLICA", "EDIFICIOS PÃšBLICOS", "Pedro Chutro", 2780,
		"", "Parque Patricios", "Comuna 4", "1437", "C1437IYD"))
	datos = append(datos, *model.CreateBicicletero(-58.4651537817402, -34.6439519627878, 781,
		"Escuela NÂ°16", 2017, "1U", 3, "VIA PÃšBLICA", "Universidades - Colegios", "Primera Junta", 3445,
		"", "Flores", "Comuna 7", "1406", "C1406HLS"))
	datos = append(datos, *model.CreateBicicletero(-58.3807476823696, -34.6046647248373, 783,
		"Ministerio de Salud", 2019, "1U", 4, "INTERIOR", "Edificio PÃºblico", "Carlos Pelegrini", 313,
		"", "San Nicolas", "Comuna 1", "1009", "C1009ABG"))
	datos = append(datos, *model.CreateBicicletero(-58.4304970607627, -34.6240541954079, 784,
		"Ochavas Valle y Doblas", 2018, "1U", -1, "VIA PÃšBLICA", "Espacio PÃºblico", "Valle", 0,
		"Doblas", "Caballito", "Comuna 6", "", ""))
	datos = append(datos, *model.CreateBicicletero(-58.4167041860391, -34.576748297251, 946,
		"Ecoparque", 2019, "1U", 7, "VIA PÃšBLICA", "Espacio PÃºblico", "SARMIENTO AV.", 2600,
		"", "", "", "", ""))
	return datos
}

var datosParseRowToBicicleteroTest = [][]string{
	{"-58.4569375558625", "-34.6266557833432", "1", "LINEA A CARABOBO", "2011", "4U", "1", "VIA PUBLICA",
		"SUBTE", "RIVADAVIA AV.", "0", "CARABOBO AV.", "Flores", "Comuna 7", "", ""},
	{"-58.3804259230237", "-34.611475939191", "107", "ASI", "2012", "4U", "1", "VIA PUBLICA",
		"EDIFICIOS PUBLICOS", "IRIGOYEN, BERNARDO DE", "272", "", "Monserrat", "Comuna 1", "1072", "C1072AAF"},
	{"-58.4263786839093", "-34.5983935322703", "457", "ESC. PRIMARIA COMUN NÂ° 20 ROSARIO VERA PEÃ‘ALOZA", "2013", "4U", "1", "VIA PUBLICA",
		"ESCUELAS", "PRINGLES", "1165", "", "Almagro", "Comuna 5", "1183", "C1183AEU"},
	{"-58.4203948587377", "-34.6547003849021", "472", "ESC. PRIMARIA COMUN NÂ° 02 DR. GENARO SISTO", "2013", "4U", "1", "VIA PUBLICA",
		"ESCUELAS", "TILCARA", "3365", "", "Nueva Pompeya", "Comuna 4", "1437", "C1437CYK"},
	{"-58.4868230292919", "-34.5527978911952", "473", "ESC. DE EDUCACION MEDIA NÂ° 05/15Â°", "2013", "4U", "1", "INTERIOR",
		"ESCUELAS", "TRONADOR", "4134", "", "Saavedra", "Comuna 12", "1430", "C1430DMZ"},
	{"-58.4084262228404", "-34.5865032086421", "478", "PLAZA LAS HERAS", "2013", "4U", "2", "VIA PUBLICA",
		"ESPACIOS VERDES", "DIAZ, CNEL. AV.", "0", "JUNCAL", "Recoleta", "Comuna 2", "", ""},
	{"-58.4660103877796", "-34.5547388764391", "701", "Catering  Schuster", "2016", "1U", "2", "VIA PUBLICA",
		"Comercio amigo", "Amenabar", "3075", "", "NuÃ±ez", "Comuna 13", "1429", "C1429AEC"},
	{"-58.410203407127", "-34.5920425550597", "702", "La MÃ¡quina de Jugar", "2016", "1U", "2", "VIA PUBLICA",
		"Comercio amigo", "SÃ¡nchez de Bustamante", "1645", "", "Recoleta", "Comuna 2", "", ""},
	{"-58.3738466790501", "-34.6026955573037", "707", "Ministerio de Ambiente de NaciÃ³n", "2016", "1U", "2", "VIA PUBLICA",
		"EDIFICIOS PÃšBLICOS", "San MartÃ­n", "451", "", "San Nicolas", "Comuna 1", "", ""},
	{"-58.3989398022133", "-34.6413855694075", "708", "Registro Nacional de las Personas RENAPER", "2016", "1U", "3", "VIA PUBLICA",
		"EDIFICIOS PÃšBLICOS", "Pedro Chutro", "2780", "", "Parque Patricios", "Comuna 4", "1437", "C1437IYD"},
	{"-58.4651537817402", "-34.6439519627878", "781", "Escuela NÂ°16", "2017", "1U", "3", "VIA PÃšBLICA",
		"Universidades - Colegios", "Primera Junta", "3445", "", "Flores", "Comuna 7", "1406", "C1406HLS"},
	{"-58.3807476823696", "-34.6046647248373", "783", "Ministerio de Salud", "2019", "1U", "4", "INTERIOR",
		"Edificio PÃºblico", "Carlos Pelegrini", "313", "", "San Nicolas", "Comuna 1", "1009", "C1009ABG"},
	{"-58.4304970607627", "-34.6240541954079", "784", "Ochavas Valle y Doblas", "2018", "1U", "None", "VIA PÃšBLICA",
		"Espacio PÃºblico", "Valle", "0", "Doblas", "Caballito", "Comuna 6", "", ""},
	{"-58.4167041860391", "-34.576748297251", "946", "Ecoparque", "2019", "1U", "7", "VIA PÃšBLICA",
		"Espacio PÃºblico", "SARMIENTO AV.", "2600", "", "", "", "", ""},
}

func TestParseRowToBicicletero(t *testing.T) {
	datosTest := createDatosParserBicicleterosTest()
	var output []model.Bicicletero
	for _, datosCadena := range datosParseRowToBicicleteroTest {
		parsed, e := parseRowToBicicletero(datosCadena)
		if e == nil {
			output = append(output, *parsed)
		}
	}
	for i := 0; i < len(datosTest); i++ {
		if datosTest[i] != output[i] {
			t.Errorf("parseRowToBicicletero: \n Encontrado:\n (%v) \n Se esperaba: \n (%v) \n", output[i], datosTest[i])
		}
	}
}

func TestParserBicicleteros(t *testing.T) {
	datosTest := createDatosParserBicicleterosTest()
	output, e := createParserBicicleterosPath(PATH_TEST_BICICLETEROS)
	if e == nil {
		for i := 0; i < len(datosTest); i++ {
			if datosTest[i] != output.bicicleteros[i] {
				t.Errorf("parserBicicleteros = \n Encontrado: \n (%v) \n Se esperaba: \n (%v) \n", output.bicicleteros[i], datosTest[i])
			}
		}
	} else {
		t.Errorf("parserBicicleteros/ nombre de archivo inválido :" + PATH_TEST_BICICLETEROS)
	}
}
