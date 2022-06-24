package model

type Bicicletero struct{
	//ubicación geográfica
	lat float32
	long float32
	//características
	id uint16
	nombre_referencia string
	anio_de_ingreso uint16
	tipo string
	cantidad_bicicleteros int16
	//dirección y clasificación
	ubicacion string
	clasificacion_lugar string
	calle string
	altura uint16
	calle_interseccion string	 
	barrio string
	comuna string
	codigo_postal string
	codigo_postal_argentino string
}

func CreateBicicletero(lat float32, long float32, id uint16, nombre_referencia string, 
					anio_de_ingreso uint16, tipo string, cantidad_bicicleteros int16,
					ubicacion string, clasificacion_lugar string,
					calle string, altura uint16, calle_interseccion string, 
					barrio string, comuna string, codigo_postal string,
					codigo_postal_argentino string) (b *Bicicletero){
	b = new (Bicicletero)
	//ubicación geográfica
	b.lat = lat
	b.long = long
	//características
	b.id = id
	b.nombre_referencia = nombre_referencia
	b.anio_de_ingreso = anio_de_ingreso
	b.tipo = tipo
	b.cantidad_bicicleteros = cantidad_bicicleteros
	//dirección y clasificación
	b.ubicacion = ubicacion
	b.clasificacion_lugar = clasificacion_lugar
	b.calle = calle
	b.altura = altura
	b.calle_interseccion = calle_interseccion
	b.barrio = barrio
	b.comuna = comuna
	b.codigo_postal = codigo_postal
	b.codigo_postal_argentino = codigo_postal_argentino
	return b
}

func (b *Bicicletero) GetLong() float32{
	return b.long
}

func (b *Bicicletero) GetLat() float32{
	return b.lat
}

func (b *Bicicletero) GetLongLat() (float32, float32){
	return b.long, b.lat
}

func (b *Bicicletero) GetId() uint16{
	return b.id
}

func (b *Bicicletero) GetNombreReferencia() string{
	return b.nombre_referencia
}

func (b *Bicicletero) GetAnioDeIngreso() uint16 {
	return b.anio_de_ingreso
}

func (b *Bicicletero) GetTipo() string{
	return b.tipo
}

func (b *Bicicletero) GetCantidadBicicleteros() int16{
	return b.cantidad_bicicleteros
}

//dirección y clasificación
func (b *Bicicletero) GetUbicacion() string{
	return b.ubicacion
}

func (b *Bicicletero) GetClasificacionLugar() string{
	return b.clasificacion_lugar
}

func (b *Bicicletero) GetCalle() string{
	return b.calle
}

func (b *Bicicletero) GetAltura() uint16{
	return b.altura
}

func (b *Bicicletero) GetCalleInterseccion() string{
	return b.calle_interseccion
} 

func (b *Bicicletero) GetBarrio() string{
	return b.barrio
}

func (b *Bicicletero) GetComuna() string{
	return b.comuna
}

func (b *Bicicletero) GetCodigoPostal() string{
	return b.codigo_postal
}

func (b *Bicicletero) GetCodigoPostalArgentino() string{
	return b.codigo_postal_argentino
}
