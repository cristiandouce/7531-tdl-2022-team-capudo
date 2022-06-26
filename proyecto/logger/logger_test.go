package logger

import (
	"os"
	"testing"
	"time"
)

type tDatoTestLog struct{
	nivel uint16 
	mensaje string
}

var datosTestLog = []tDatoTestLog{
	{nivel: DEBUG, mensaje: "Item 1 debugging"},
	{nivel: INFO, mensaje: "Item 1 info"},
	{nivel: WARNING, mensaje: "Item 1 warning"},
	{nivel: ERROR, mensaje: "Item 1 error"},
	{nivel: CRITICAL, mensaje: "Item 1 critical"},
	{nivel: DEBUG, mensaje: "Item 2 debugging"},
	{nivel: INFO, mensaje: "Item 2 info"},
	{nivel: ERROR, mensaje: "Item 2 error"},
	{nivel: WARNING, mensaje: "Item 2 warning"},
	{nivel: CRITICAL, mensaje: "Item 2 critical"},
	{nivel: WARNING, mensaje: "Item 3 warning"},
	{nivel: INFO, mensaje: "Item 3 info"},
	{nivel: ERROR, mensaje: "Item 3 error"},
	{nivel: INFO, mensaje: "Item 4 info"},
	{nivel: DEBUG, mensaje: "Item 3 debugging"},
	{nivel: WARNING, mensaje: "Item 4 warning"},
	{nivel: CRITICAL, mensaje: "Item 3 critical"},
	{nivel: CRITICAL, mensaje: "Item 4 critical"},
	{nivel: INFO, mensaje: "Item 5 info"},
	{nivel: ERROR, mensaje: "Item 4 error"},
	{nivel: DEBUG, mensaje: "Item 4 debugging"},
	{nivel: WARNING, mensaje: "Item 5 warning"},
	{nivel: INFO, mensaje: "Item 6 info"},
	{nivel: INFO, mensaje: "Item 7 info"},
	{nivel: ERROR, mensaje: "Item 5 error"},
	{nivel: WARNING, mensaje: "Item 6 warning"},
	{nivel: CRITICAL, mensaje: "Item 5 critical"},
	{nivel: ERROR, mensaje: "Item 6 error"},
	{nivel: DEBUG, mensaje: "Item 5 debugging"},
	{nivel: DEBUG, mensaje: "Item 6 debugging"},
	{nivel: WARNING, mensaje: "Item 7 warning"},
	{nivel: CRITICAL, mensaje: "Item 6 critical"},
	{nivel: CRITICAL, mensaje: "Item 7 critical"},
	{nivel: CRITICAL, mensaje: "Item 8 critical"},
	{nivel: ERROR, mensaje: "Item 7 error"},
	{nivel: INFO, mensaje: "Item 8 info"},
	{nivel: INFO, mensaje: "Item 9 info"},
	{nivel: INFO, mensaje: "Item 10 info"},
	{nivel: WARNING, mensaje: "Item 8 warning"},
	{nivel: DEBUG, mensaje: "Item 7 debugging"},
	{nivel: ERROR, mensaje: "Item 8 error"},
	{nivel: WARNING, mensaje: "Item 9 warning"},
	{nivel: CRITICAL, mensaje: "Item 9 critical"},
	{nivel: DEBUG, mensaje: "Item 8 debugging"},
	{nivel: INFO, mensaje: "Item  info"},
	{nivel: DEBUG, mensaje: "Item 9 debugging"},
	{nivel: ERROR, mensaje: "Item 9 error"},
	{nivel: INFO, mensaje: "Item 11 info"},
	{nivel: WARNING, mensaje: "Item 10 warning"},
	{nivel: CRITICAL, mensaje: "Item 10 critical"},
	{nivel: CRITICAL, mensaje: "Item 11 critical"},
	{nivel: DEBUG, mensaje: "Item 10 debugging"},
	{nivel: INFO, mensaje: "Item 12 info"},
	{nivel: INFO, mensaje: "Item 13 info"},
	{nivel: WARNING, mensaje: "Item 11 warning"},
	{nivel: ERROR, mensaje: "Item 10 error"},
	{nivel: WARNING, mensaje: "Item 12 warning"},
	{nivel: ERROR, mensaje: "Item 11 error"},
	{nivel: CRITICAL, mensaje: "Item 12 critical"},
	{nivel: CRITICAL, mensaje: "Item 13 critical"},
}

func tryLogDato(dato tDatoTestLog) (e error){
	switch dato.nivel {
		case DEBUG:
			e = LogDebug(dato.mensaje)
		case INFO :
			e = LogInfo(dato.mensaje)
		case WARNING :
			e = LogWarning(dato.mensaje)
		case ERROR : 
			e = LogError(dato.mensaje)
		case CRITICAL :
			e = LogCritical(dato.mensaje)
	}
	return e
}

func TestLoggerDefault(t *testing.T) {
	os.Remove(PATH_LOG)
	for _, dato := range datosTestLog {
		e := tryLogDato(dato)
		if(e != nil){
			t.Errorf("Logger Test Default = (%v), se esperaba (%v)",e.Error(), "nil")
		}
	}
}

func TestLoggerInfo(t *testing.T) {
	SetNivelLogger(INFO)
	for _, dato := range datosTestLog {
		e := tryLogDato(dato)
		if((e != nil) && (dato.nivel >= INFO)){
			t.Errorf("Logger Test Info = (%v), se esperaba (%v)",e.Error(), "nil")
		}
	}
}

func TestLoggerWarning(t *testing.T) {
	SetNivelLogger(WARNING)
	for _, dato := range datosTestLog {
		e := tryLogDato(dato)
		if((e != nil) && (dato.nivel >= WARNING)){
			t.Errorf("Logger Test Warning = (%v), se esperaba (%v)",e.Error(), "nil")
		}
	}
}

func TestLoggerError(t *testing.T) {
	SetNivelLogger(ERROR)
	for _, dato := range datosTestLog {
		e := tryLogDato(dato)
		if((e != nil) && (dato.nivel >= ERROR)){
			t.Errorf("Logger Test Error = (%v), se esperaba (%v)",e.Error(), "nil")
		}
	}
}

func TestLoggerCritical(t *testing.T) {
	SetNivelLogger(CRITICAL)
	for _, dato := range datosTestLog {
		e := tryLogDato(dato)
		if((e != nil) && (dato.nivel >= CRITICAL)){
			t.Errorf("Logger Test Critical = (%v), se esperaba (%v)",e.Error(), "nil")
		}
	}
}

func TestLoggerConcurrente(t *testing.T) {
	SetNivelLogger(DEBUG)
	go func() {
		index1 := 0	
		for {
			if (index1 >= len(datosTestLog)){
				break
			} else{
				dato := datosTestLog[index1]
				dato.mensaje += " - Hilo Secundario"
				e := tryLogDato(dato)			
				if(e != nil){
					t.Errorf("Logger Test Concurrente = (%v), se esperaba (%v)",e.Error(), "nil")
				}
				index1++
			}	
			time.Sleep(time.Millisecond * time.Duration(5))
		}	
	}()
	index2 := 0
	for {
		if (index2 >= len(datosTestLog)){
			break
		} else{
			dato := datosTestLog[index2]
			dato.mensaje += " - Hilo Padre"
			e := tryLogDato(dato)			
			if(e != nil){
				t.Errorf("Logger Test Concurrente = (%v), se esperaba (%v)",e.Error(), "nil")
			}
			index2++
		}			
		time.Sleep(time.Millisecond * time.Duration(5))
	}	
}
