package model

import (
	"time"
)

type Recorrido struct{
	//id y duración
	id_recorrido string
	duraccion_recorrido uint32
	//origen
	fecha_origen_recorrido time.Time
	id_estacion_origen string
	nombre_estacion_origen string
	direccion_estacion_origen string
	long_estacion_origen float32
	lat_estacion_origen float32
	//destino
	fecha_destino_recorrido time.Time
	id_estacion_destino string
	nombre_estacion_destino string
	direccion_estacion_destino string
	long_estacion_destino float32
	lat_estacion_destino float32
	//datos usuario y bici
	id_usuario string
	modelo_bicicleta string
}

func CreateRecorrido(id_recorrido string, duraccion_recorrido uint32,
			fecha_origen_recorrido time.Time, id_estacion_origen string,
			nombre_estacion_origen string, direccion_estacion_origen string,
			long_estacion_origen float32, lat_estacion_origen float32,
			fecha_destino_recorrido time.Time, id_estacion_destino string,
			nombre_estacion_destino string, direccion_estacion_destino string,
			long_estacion_destino float32, lat_estacion_destino float32,
			id_usuario string, modelo_bicicleta string) (r *Recorrido){
	r = new (Recorrido)
	//id y duración
	r.id_recorrido = id_recorrido
	r.duraccion_recorrido = duraccion_recorrido
	//origen
	r.fecha_origen_recorrido = fecha_origen_recorrido
	r.id_estacion_origen = id_estacion_origen
	r.nombre_estacion_origen = nombre_estacion_origen
	r.direccion_estacion_origen = direccion_estacion_origen
	r.long_estacion_origen = long_estacion_origen
	r.lat_estacion_origen = lat_estacion_origen
	//destino
	r.fecha_destino_recorrido = fecha_destino_recorrido
	r.id_estacion_destino = id_estacion_destino
	r.nombre_estacion_destino = nombre_estacion_destino
	r.direccion_estacion_destino = direccion_estacion_destino
	r.long_estacion_destino = long_estacion_destino
	r.lat_estacion_destino = lat_estacion_destino
	//datos usuario y bicicleta
	r.id_usuario = id_usuario
	r.modelo_bicicleta = modelo_bicicleta
	return r
}
