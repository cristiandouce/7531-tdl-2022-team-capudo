package parser

import (
	"capudo/model"
	"testing"
)

const PATH_TEST_USUARIOS string = "test_usuarios.csv"

func createDatosParserUsuarioTest() (datos []model.Usuario) {
	datos = append(datos, *model.CreateUsuario(605638, "OTRO",
		27, createTiempo("2020-09-29 13:03:00 UTC")))
	datos = append(datos, *model.CreateUsuario(605550, "F",
		18, createTiempo("2020-09-21 18:03:00 UTC")))
	datos = append(datos, *model.CreateUsuario(605347, "F",
		37, createTiempo("2020-02-27 11:23:00 UTC")))
	datos = append(datos, *model.CreateUsuario(605130, "F",
		25, createTiempo("2020-09-10 14:57:00 UTC")))
	datos = append(datos, *model.CreateUsuario(605119, "F",
		35, createTiempo("2020-02-27 11:23:00 UTC")))
	datos = append(datos, *model.CreateUsuario(604773, "F",
		50, createTiempo("2020-09-25 0:11:00 UTC")))

	return datos
}

var datosParseRowToUsuarioTest = [][]string{
	//{"id_usuario","genero_usuario","edad_usuario","fecha_alta"}
	{"605638", "OTRO", "27", "29-09-20", "1:03:00 PM"},
	{"605550", "F", "18", "21-09-20", "6:03:00 PM"},
	{"605347", "F", "37", "27-02-20", "11:23:00 AM"},
	{"605130", "F", "25", "10-09-20", "2:57:00 PM"},
	{"605119", "F", "35", "27-02-20", "11:23:00 AM"},
	{"604773", "F", "50", "25-09-20", "12:11:00 AM"},
}

/*
func createTiempo(cadena string)(t time.Time){
	t, _ = ParseTimeUTC(cadena)
	return t
}*/

func TestParseRowToUsuario(t *testing.T) {
	datosTest := createDatosParserUsuarioTest()
	var output []model.Usuario
	for _, datosCadena := range datosParseRowToUsuarioTest {
		parsed, e := parseRowToUsuario(datosCadena)
		if e == nil {
			output = append(output, *parsed)
		}
	}
	for i := 0; i < len(datosTest); i++ {
		if datosTest[i] != output[i] {
			t.Errorf("parseRowToUsuario = (%v), se esperaba (%v)", output[i], datosTest[i])
		}
	}
}

func TestParserUsuarios(t *testing.T) {
	datosTest := createDatosParserUsuarioTest()
	output, e := createParserUsuariosPath(PATH_TEST_USUARIOS)
	if e == nil {
		for i := 0; i < len(output.usuarios); i++ {
			if datosTest[i] != output.usuarios[i] {
				t.Errorf("parserRecorridos = (%v), se esperaba (%v)", output.usuarios[i], datosTest[i])
			}
		}
	} else {
		t.Errorf("parserUsuarios/ nombre de archivo inválido :" + PATH_TEST_USUARIOS)
	}
}
