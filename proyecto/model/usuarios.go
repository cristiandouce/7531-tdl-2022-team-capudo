package model

import (
	"encoding/json"
	"time"
)

type Usuario struct {
	//id
	id_usuario uint32
	//datos usuario
	genero_usuario string
	edad_usuario   uint16
	//fecha de alta
	fecha_alta time.Time
}

func CreateUsuario(id_usuario uint32, genero_usuario string, edad_usuario uint16,
	fecha_alta time.Time) (u *Usuario) {
	u = new(Usuario)
	//id
	u.id_usuario = id_usuario
	//datos usuario
	u.genero_usuario = genero_usuario
	u.edad_usuario = edad_usuario
	//fecha alta
	u.fecha_alta = fecha_alta
	return u
}

func (u *Usuario) GetIdUsuario() uint32 {
	return u.id_usuario
}
func (u *Usuario) GetGeneroUsuario() string {
	return u.genero_usuario
}
func (u *Usuario) GetEdadUsuario() uint16 {
	return u.edad_usuario
}
func (u *Usuario) GetFechaAlta() time.Time {
	return u.fecha_alta
}

// Ejemplo de serialización de miembros privados al convertir a JSON
// especificamente cuando se usa `encoder/json`
func (u *Usuario) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID         uint32    `json:"id"`
		GENERO     string    `json:"genero"`
		EDAD       uint16    `json:"edad"`
		FECHA_ALTA time.Time `json:"fecha_alta"`
	}{
		ID:         u.id_usuario,
		GENERO:     u.genero_usuario,
		EDAD:       u.edad_usuario,
		FECHA_ALTA: u.fecha_alta})
}
