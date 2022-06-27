package logger

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

const PATH_LOG = "/tmp/log.txt"
const DEBUG uint16 = 0
const INFO uint16 = 1
const WARNING uint16 = 2
const ERROR uint16 = 3
const CRITICAL uint16 = 4
const DEFAULT uint16 = DEBUG

var lock_logger = &sync.Mutex{}

// Instancia del singleton del logger
var (
	instanceLog *logger
)

type logger struct {
	nivel   uint16
	archivo *os.File
}

func createLog() (e error) {
	instanceLog = new(logger)
	instanceLog.archivo, e = os.OpenFile(PATH_LOG, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	instanceLog.nivel = DEFAULT
	return e
}

func checkInstanceLogAndCreateIfNil() (e error) {
	e = nil
	if instanceLog == nil {
		e = createLog()
	}
	return e
}

func addLabelAndTimeToMessage(param string, nivel uint16) (retorno string) {
	switch nivel {
	case DEBUG:
		retorno = "DEBUG"
	case INFO:
		retorno = "INFO"
	case WARNING:
		retorno = "WARNING"
	case ERROR:
		retorno = "ERROR"
	case CRITICAL:
		retorno = "CRITICAL"
	}
	retorno += " [" + time.Now().String() + "]: " + param + "\n"
	return retorno
}

func checkNivelAndDoLog(nivel uint16, data string) (e error) {
	e = nil
	if nivel >= instanceLog.nivel {
		_, e = instanceLog.archivo.WriteString(addLabelAndTimeToMessage(data, nivel))
		if instanceLog.nivel == DEBUG {
			fmt.Print(addLabelAndTimeToMessage(data, nivel))
		}
	} else {
		e = fmt.Errorf("Error no registrable - Nivel Actual:" + strconv.Itoa(int(instanceLog.nivel)))
	}
	return e
}

func doLog(nivel uint16, data string) (e error) {
	lock_logger.Lock()
	defer lock_logger.Unlock()
	e = checkInstanceLogAndCreateIfNil()
	if e == nil {
		e = checkNivelAndDoLog(nivel, data)
	}
	return e
}

/*
 	Información para el proceso de debugging.
	Ejemplo: Contenido de una consulta y su respuesta en forma detallada.
*/
func LogDebug(data string) (e error) {
	return doLog(DEBUG, data)
}

/*
 	Confirmación de eventos durante el funcionamiento correcto.
	Ejemplo: Se cargaron los archivos con la base de datos o registro de consultas.
*/
func LogInfo(data string) (e error) {
	return doLog(INFO, data)
}

/* 	Errores que permiten que el programa sigan funcionando correctamente.
Ejemplo: No se pudo cargar correctamente alguno de los registros de la base de datos.
*/
func LogWarning(data string) (e error) {
	return doLog(WARNING, data)
}

/*
	Errores que inhabilitan una funcionalidad completa.
	Ejemplo: No se pudo cargar alguno de los archivos de la base de datos.
*/
func LogError(data string) (e error) {
	return doLog(ERROR, data)
}

//	Errores que no permiten que el programa siga funcionando.
func LogCritical(data string) (e error) {
	return doLog(CRITICAL, data)
}

// Permite setear el nivel de log
func SetNivelLogger(nivel uint16) (e error) {
	lock_logger.Lock()
	defer lock_logger.Unlock()
	e = checkInstanceLogAndCreateIfNil()
	if e == nil {
		instanceLog.nivel = nivel
	}
	return e
}
