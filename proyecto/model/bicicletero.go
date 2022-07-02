package model

import "strconv"

type FeatureBicicletero struct {
	Tipo       string              `json:"type"`
	Properties EstacionBicicletero `json:"properties"`
	Geometry   Ubicacion           `json:"geometry"`
}
type EstacionBicicletero struct {
	Id        uint32 `json:"id"`
	Codigo    uint32 `json:"codigo"`
	Nombre    string `json:"nombre"`
	Ubicacion string `json:"ubicacion"`
	Tipo      string `json:"tipo"`
	Horario   string `json:"horario"`
	Anclajes  uint16 `json:"anclajes_t"`
}

type Ubicacion struct {
	Tipo        string     `json:"tipo"`
	Coordinates [2]float32 `json:"coordinates"`
}
type Bicicletero struct {
	//características
	Id         uint32 `json:"id"`
	Codigo     string `json:"nombre_referencia"`
	Nombre     string `json:"nombre_estacion"`
	Ubicacion  string `json:"ubicacion"`
	Tipo       string `json:"tipo"`
	Horario    string `json:"horario"`
	Anclajes_t uint16 `json:"anclajes_t"`
	//ubicación geográfica
	Lat  float32 `json:"latitud"`
	Long float32 `json:"longitud"`
}

func ParserFeatureToBicicletero(feature FeatureBicicletero) Bicicletero {
	var b Bicicletero
	b.Id = feature.Properties.Id
	b.Codigo = strconv.FormatUint(uint64(feature.Properties.Codigo), 10) + "BAEcobici"
	b.Nombre = feature.Properties.Nombre
	b.Ubicacion = feature.Properties.Ubicacion
	b.Tipo = feature.Properties.Tipo
	b.Horario = feature.Properties.Horario
	b.Anclajes_t = feature.Properties.Anclajes
	b.Long = feature.Geometry.Coordinates[0]
	b.Lat = feature.Geometry.Coordinates[1]
	return b
}
