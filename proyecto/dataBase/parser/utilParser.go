package parser

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

/*
	Params: cadena (flotante como string), eParam (error arrastrado).
	Returns: fotante32 (flotante resultado de la conversión de la cadena),
		error(se retorna el primer error en aparición entre el parámetro y el error
		generado en la conversión)
*/
func ParseToFloat32Error(cadena string, eParam error) (flotante32 float32, e error) {
	flotante64, e := strconv.ParseFloat(cadena, 32)
	flotante32 = float32(flotante64)
	if (eParam != nil){
		e = eParam
	}	
	return flotante32, e
}

/* Params: cadena (entero como cadena), eParam (error)
   Returns: entero (entero convertido desde la cadena), 
   			e (el primer error en orden de aparación entre eParam y 
			el error de conversión)
*/
func ParseToUint32Error(cadena string, eParam error) (entero uint32, e error){
	var entero64 uint64
	entero64, e = strconv.ParseUint(cadena, 10, 32)
	entero = uint32(entero64)
	if (eParam != nil){
		e = eParam
	}
	return entero, e
}

/* Params: cadena (entero como cadena), eParam (error)
   Returns: entero (entero convertido desde la cadena), 
   			e (el primer error en orden de aparación entre eParam y 
			el error de conversión)
*/
func ParseToUint16Error(cadena string, eParam error) (entero uint16, e error){
	var entero64 uint64
	entero64, e = strconv.ParseUint(cadena, 10, 16)
	entero = uint16(entero64)
	if (eParam != nil){
		e = eParam
	}
	return entero, e
}

/* 
	Params: cadena (entero como cadena), eParam (error)
   	Returns: entero (entero convertido desde la cadena), 
   			e (el primer error en orden de aparición entre eParam y 
			el error de conversión en esta función)
*/
func ParseToIntError(cadena string, eParam error) (entero int, e error){
	entero, e = strconv.Atoi(cadena)
	if (eParam != nil){
		e = eParam
	}
	return entero, e
}

/*
	Params: tiempoCadena (fecha y hora en formato UTC como cadena),
	Returns: t (la fecha convertida al tipo Time),
			e (el primer error al intentar convertir la fecha)
*/
func ParseTimeUTC(tiempoCadena string) (t time.Time, e error){
	vectorTiempo := strings.Split(tiempoCadena, " ")
	// obtención componentes de fecha
	var dia, mes, anio int
	vectorFecha := strings.Split(vectorTiempo[0], "-")
	if (len(vectorFecha) >= 3){
		dia, e = ParseToIntError(vectorFecha[2], e)
		mes, e = ParseToIntError(vectorFecha[1], e)
		anio, e = ParseToIntError(vectorFecha[0], e)
	} else {
		e = fmt.Errorf("formato incorrecto de fecha")
	}	
	// obtención componentes de la hora
	var hora, minuto, segundos int
	vectorHora := strings.Split(vectorTiempo[1], ":")
	if (len(vectorHora) >= 3){
		hora, e = ParseToIntError(vectorHora[0], e)
		minuto, e = ParseToIntError(vectorHora[1], e)
		segundos, e = ParseToIntError(vectorHora[2], e)
	} else if (e == nil)  {
		e = fmt.Errorf("formato incorrecto de hora")
	}
	t = time.Date(anio, time.Month(mes), dia, hora, minuto, segundos, 0, time.UTC)
	return t, e
}

/*
	Params: tiempoCadena (fecha y hora en formato UTC como cadena),
	Returns: t (la fecha convertida al tipo Time),
			e (el primer error al intentar convertir la fecha)
*/
func ParseTimeUTCError(tiempoCadena string, eParam error) (t time.Time, e error){
	t, e = ParseTimeUTC(tiempoCadena)
	if (eParam != nil){
		e = eParam
	}
	return t, e
}

/* Params: cadena (entero 16 bits como cadena), eParam (error)
   Returns: entero (entero convertido desde la cadena), 
   			e (el primer error en orden de aparación entre eParam y 
			el error de conversión)
			En caso de cadena vacía o con valor None se retorna -1 sin retornar error
			propio de esta función.
*/
func ParseToInt16WithEmptyError(cadenaParam string, eParam error) (entero16 int16, e error){
	entero16 = -1
	if((len(cadenaParam) > 0) && (!strings.Contains(cadenaParam,"None"))){
		var enteroU16 uint16
		enteroU16, e = ParseToUint16Error(cadenaParam, eParam)
		entero16 = int16(enteroU16)			
	} else {
		e = eParam
	}
	return entero16, e
}