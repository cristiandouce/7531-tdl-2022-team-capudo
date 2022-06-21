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