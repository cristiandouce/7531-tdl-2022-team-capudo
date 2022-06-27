package model

type Bicicletero struct {
	//ubicación geográfica
	Lat  float32 `json:"latitud"`
	Long float32 `json:"longitud"`
	//características
	ID                   uint16 `json:"id"`
	NombreReferencia     string `json:"nombre_referencia"`
	AnioIngreso          uint16 `json:"anio_ingreso"`
	Tipo                 string `json:"tipo"`
	CantidadBicicleteros int16  `json:"cantidad_bicicleteros"`
	//dirección y clasificación
	Ubicacion             string `json:"ubicacion"`
	ClasificacionLugar    string `json:"clasificacion_lugar"`
	Calle                 string `json:"calle"`
	Altura                uint16 `json:"altura"`
	CalleInterseccion     string `json:"calle_interseccion"`
	Barrio                string `json:"barrio"`
	Comuna                string `json:"comuna"`
	CodigoPostal          string `json:"codigo_postal"`
	CodigoPostalArgentino string `json:"codigo_postal_argentino"`
}

func CreateBicicletero(Lat float32, Long float32, ID uint16, NombreReferencia string,
	AnioIngreso uint16, Tipo string, CantidadBicicleteros int16,
	Ubicacion string, ClasificacionLugar string,
	Calle string, Altura uint16, CalleInterseccion string,
	Barrio string, Comuna string, CodigoPostal string,
	CodigoPostalArgentino string) (b *Bicicletero) {
	b = new(Bicicletero)
	//ubicación geográfica
	b.Lat = Lat
	b.Long = Long
	//características
	b.ID = ID
	b.NombreReferencia = NombreReferencia
	b.AnioIngreso = AnioIngreso
	b.Tipo = Tipo
	b.CantidadBicicleteros = CantidadBicicleteros
	//dirección y clasificación
	b.Ubicacion = Ubicacion
	b.ClasificacionLugar = ClasificacionLugar
	b.Calle = Calle
	b.Altura = Altura
	b.CalleInterseccion = CalleInterseccion
	b.Barrio = Barrio
	b.Comuna = Comuna
	b.CodigoPostal = CodigoPostal
	b.CodigoPostalArgentino = CodigoPostalArgentino
	return b
}

func (b *Bicicletero) GetLong() float32 {
	return b.Long
}

func (b *Bicicletero) GetLat() float32 {
	return b.Lat
}

func (b *Bicicletero) GetLongLat() (float32, float32) {
	return b.Long, b.Lat
}

func (b *Bicicletero) GetId() uint16 {
	return b.ID
}

func (b *Bicicletero) GetNombreReferencia() string {
	return b.NombreReferencia
}

func (b *Bicicletero) GetAnioDeIngreso() uint16 {
	return b.AnioIngreso
}

func (b *Bicicletero) GetTipo() string {
	return b.Tipo
}

func (b *Bicicletero) GetCantidadBicicleteros() int16 {
	return b.CantidadBicicleteros
}

//dirección y clasificación
func (b *Bicicletero) GetUbicacion() string {
	return b.Ubicacion
}

func (b *Bicicletero) GetClasificacionLugar() string {
	return b.ClasificacionLugar
}

func (b *Bicicletero) GetCalle() string {
	return b.Calle
}

func (b *Bicicletero) GetAltura() uint16 {
	return b.Altura
}

func (b *Bicicletero) GetCalleInterseccion() string {
	return b.CalleInterseccion
}

func (b *Bicicletero) GetBarrio() string {
	return b.Barrio
}

func (b *Bicicletero) GetComuna() string {
	return b.Comuna
}

func (b *Bicicletero) GetCodigoPostal() string {
	return b.CodigoPostal
}

func (b *Bicicletero) GetCodigoPostalArgentino() string {
	return b.CodigoPostalArgentino
}
