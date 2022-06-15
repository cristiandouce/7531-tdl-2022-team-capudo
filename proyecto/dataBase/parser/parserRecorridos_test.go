package parser

import (
	"capudo/model"
	"testing"
	"time"
)

const PATH_TEST_RECORRIDOS string = "test_trips_2021.csv"

func createDatosParserRecorridoTest()(datos []model.Recorrido){
	datos = append(datos, *model.CreateRecorrido("10758473BAEcobici", 591, 
		createTiempo("2021-04-10 20:38:24 UTC"), "2BAEcobici", "002 - Retiro I", 
		"Ramos Mejia, Jose Maria, Dr. Av. & Del Libertador Av.",
		-58.3747151, -34.5924233, 
		createTiempo("2021-04-10 20:48:15 UTC"), "95BAEcobici", "095 - ESMERALDA", 
		"ESMERALDA 516",
		-58.3781678, -34.6021121,
		"86840BAEcobici", "ICONIC"))
	datos = append(datos, *model.CreateRecorrido("10757803BAEcobici",1321,
		createTiempo("2021-04-10 16:34:08 UTC"), "2BAEcobici", "002 - Retiro I",
		"Ramos Mejia, Jose Maria, Dr. Av. & Del Libertador Av.",
		-58.3747151, -34.5924233,
		createTiempo("2021-04-10 16:56:09 UTC"), "73BAEcobici", "073 - Ruy Díaz de Guzmán",
		"Avenida Martin Garcia y Ruy Díaz de Guzmán",
		-58.3718235, -34.6306814,
		"52860BAEcobici","ICONIC"))
	datos = append(datos, *model.CreateRecorrido("10756603BAEcobici", 380,
		createTiempo("2021-04-10 07:06:00 UTC"), "3BAEcobici", "003 - ADUANA",
		"Moreno & Av Paseo Colon", 
		-58.3682604, -34.611032,
		createTiempo("2021-04-10 07:12:20 UTC"), "150BAEcobici","150 - RODRIGO BUENO",
		"Av. España 2200",
		-58.3554654, -34.6187547,
		"375594BAEcobici","ICONIC"))	
	datos = append(datos, *model.CreateRecorrido("10756618BAEcobici", 1436,
		createTiempo("2021-04-10 07:25:08 UTC"), "4BAEcobici", "004 - Plaza Roma",
		"Lavalle & Bouchard", 
		-58.368781, -34.601822, 
		createTiempo("2021-04-10 07:49:04 UTC"), "353BAEcobici", "237 - Madero Office",
		"367 Sanchez De Thompson, Mariquita",
		-58.364695, -34.599036,
		"489972BAEcobici", "ICONIC"))
	return datos
}

var datosParseRowToRecorridoTest = [][] string{
//	{"id_recorrido", "duracion_recorrido", "fecha_origen_recorrido", "id_estacion_origen", "nombre_estacion_origen", "direccion_estacion_origen", "long_estacion_origen", "lat_estacion_origen", "fecha_destino_recorrido", "id_estacion_destino", "nombre_estacion_destino", "direccion_estacion_destino", "long_estacion_destino", "lat_estacion_destino", "id_usuario", "modelo_bicicleta"},
	{"10758473BAEcobici","591","2021-04-10 20:38:24 UTC","2BAEcobici","002 - Retiro I","Ramos Mejia, Jose Maria, Dr. Av. & Del Libertador Av.","-58.3747151","-34.5924233","2021-04-10 20:48:15 UTC","95BAEcobici","095 - ESMERALDA","ESMERALDA 516","-58.3781678","-34.6021121","86840BAEcobici","ICONIC"},
	{"10757803BAEcobici","1321","2021-04-10 16:34:08 UTC","2BAEcobici","002 - Retiro I","Ramos Mejia, Jose Maria, Dr. Av. & Del Libertador Av.","-58.3747151","-34.5924233","2021-04-10 16:56:09 UTC","73BAEcobici","073 - Ruy Díaz de Guzmán","Avenida Martin Garcia y Ruy Díaz de Guzmán","-58.3718235","-34.6306814","52860BAEcobici","ICONIC"},
	{"10756603BAEcobici","380","2021-04-10 07:06:00 UTC","3BAEcobici","003 - ADUANA","Moreno & Av Paseo Colon","-58.3682604","-34.611032","2021-04-10 07:12:20 UTC","150BAEcobici","150 - RODRIGO BUENO","Av. España 2200","-58.3554654","-34.6187547","375594BAEcobici","ICONIC"},
	{"10756618BAEcobici","1436","2021-04-10 07:25:08 UTC","4BAEcobici","004 - Plaza Roma","Lavalle & Bouchard","-58.368781","-34.601822","2021-04-10 07:49:04 UTC","353BAEcobici","237 - Madero Office","367 Sanchez De Thompson, Mariquita","-58.364695","-34.599036","489972BAEcobici","ICONIC"},
}

func createTiempo(cadena string)(t time.Time){
	t, _ = ParseTimeUTC(cadena)
	return t
}

func TestParseRowToRecorrido(t *testing.T){
	datosTest := createDatosParserRecorridoTest()
	var output []model.Recorrido
	for _, datosCadena := range datosParseRowToRecorridoTest{
		parsed, e := parseRowToRecorrido(datosCadena)
		if (e == nil){
			output = append(output, *parsed)
		}
	} 
	for i := 0; i < 4; i++  {
		if(datosTest[i] != output[i]){
			t.Errorf("parseRowToRecorrido = (%v), se esperaba (%v)",output[i],datosTest[i])
		}
	}
}

func TestParserRecorridos(t *testing.T) {
	datosTest := createDatosParserRecorridoTest()
	output, e:= createParserRecorridosPath(PATH_TEST_RECORRIDOS)
	if (e == nil){
		for i := 0; i < 4; i++  {
			if(datosTest[i] != output.recorridos[i]){
				t.Errorf("parserRecorridos = (%v), se esperaba (%v)",output.recorridos[i],datosTest[i])
			}
		}
	} else {
		t.Errorf("parserRecorridos/ nombre de archivo inválido :"+ PATH_TEST_RECORRIDOS)
	}
}
