package model

import (
	"time"
)

type Recorrido struct {
	//id y duración
	IDRecorrido       string `json:"id"`
	DuracionRecorrido uint32 `json:"duracion_recorrido"`
	//origen
	FechaOrigenRecorrido    time.Time `json:"fecha_origen_recorrido"`
	IDEstacionOrigen        string    `json:"id_estacion_origen"`
	NombreEstacionOrigen    string    `json:"nombre_estacion_origen"`
	DireccionEstacionOrigen string    `json:"direccion_estacion_origen"`
	LongEstacionOrigen      float32   `json:"longitud_estacion_origen"`
	LatEstacionOrigen       float32   `json:"latitud_estacion_origen"`
	//destino
	FechaDestinoRecorrido    time.Time `json:"fecha_destino_recorrido"`
	IDEstacionDestino        string    `json:"id_estacion_destino"`
	NombreEstacionDestino    string    `json:"nombre_estacion_destino"`
	DireccionEstacionDestino string    `json:"direccion_estacion_destino"`
	LongEstacionDestino      float32   `json:"longitud_estacion_destino"`
	LatEstacionDestino       float32   `json:"latitud_estacion_destino"`
	//datos usuario y bici
	IDUsuario       string `json:"id_usuario"`
	ModeloBicicleta string `json:"modelo_bicicleta"`
}

func CreateRecorrido(IDRecorrido string, DuracionRecorrido uint32,
	FechaOrigenRecorrido time.Time, IDEstacionOrigen string,
	NombreEstacionOrigen string, DireccionEstacionOrigen string,
	LongEstacionOrigen float32, LatEstacionOrigen float32,
	FechaDestinoRecorrido time.Time, IDEstacionDestino string,
	NombreEstacionDestino string, DireccionEstacionDestino string,
	LongEstacionDestino float32, LatEstacionDestino float32,
	IDUsuario string, ModeloBicicleta string) (r *Recorrido) {
	r = new(Recorrido)
	//id y duración
	r.IDRecorrido = IDRecorrido
	r.DuracionRecorrido = DuracionRecorrido
	//origen
	r.FechaOrigenRecorrido = FechaOrigenRecorrido
	r.IDEstacionOrigen = IDEstacionOrigen
	r.NombreEstacionOrigen = NombreEstacionOrigen
	r.DireccionEstacionOrigen = DireccionEstacionOrigen
	r.LongEstacionOrigen = LongEstacionOrigen
	r.LatEstacionOrigen = LatEstacionOrigen
	//destino
	r.FechaDestinoRecorrido = FechaDestinoRecorrido
	r.IDEstacionDestino = IDEstacionDestino
	r.NombreEstacionDestino = NombreEstacionDestino
	r.DireccionEstacionDestino = DireccionEstacionDestino
	r.LongEstacionDestino = LongEstacionDestino
	r.LatEstacionDestino = LatEstacionDestino
	//datos usuario y bicicleta
	r.IDUsuario = IDUsuario
	r.ModeloBicicleta = ModeloBicicleta
	return r
}

func (r *Recorrido) GetIdRecorrido() string {
	return r.IDRecorrido
}

func (r *Recorrido) GetDuraccionRecorrido() uint32 {
	return r.DuracionRecorrido
}

//origen
func (r *Recorrido) GetFechaOrigenRecorrido() time.Time {
	return r.FechaOrigenRecorrido
}

//origen
func (r *Recorrido) GetIdEstacionOrigen() string {
	return r.IDEstacionOrigen
}

//origen
func (r *Recorrido) GetNombreEstacionOrigen() string {
	return r.NombreEstacionOrigen
}

//origen
func (r *Recorrido) GetDireccionEstacionOrigen() string {
	return r.DireccionEstacionOrigen
}

//origen
func (r *Recorrido) GetLongEstacionOrigen() float32 {
	return r.LongEstacionOrigen
}

//origen
func (r *Recorrido) GetLatEstacionOrigen() float32 {
	return r.LatEstacionOrigen
}

//origen
func (r *Recorrido) GetLongLatEstacionOrigen() (float32, float32) {
	return r.LongEstacionOrigen, r.LatEstacionOrigen
}

//destino
func (r *Recorrido) GetFechaDestinoRecorrido() time.Time {
	return r.FechaDestinoRecorrido
}

//destino
func (r *Recorrido) GetIdEstacionDestino() string {
	return r.IDEstacionDestino
}

//destino
func (r *Recorrido) GetNombreEstacionDestino() string {
	return r.NombreEstacionDestino
}

//destino
func (r *Recorrido) GetDireccionEstacionDestino() string {
	return r.DireccionEstacionDestino
}

//destino
func (r *Recorrido) GetLongEstacionDestino() float32 {
	return r.LongEstacionDestino
}

//destino
func (r *Recorrido) GetLatEstacionDestino() float32 {
	return r.LatEstacionDestino
}

//destino
func (r *Recorrido) GetLongLatEstacionDestino() (float32, float32) {
	return r.LongEstacionDestino, r.LatEstacionDestino
}

func (r *Recorrido) GetIdUsuario() string {
	return r.IDUsuario
}

func (r *Recorrido) GetModeloBicicleta() string {
	return r.ModeloBicicleta
}
