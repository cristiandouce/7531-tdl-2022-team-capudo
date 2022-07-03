# REST API en Go - para consultas de ECOBICI

> Desarrollado por el **TEAM CAPUDO** para la materia **7531 - Teoría del Lenguaje* de la [Facultad de Ingeniería de la UBA](https://fi.uba.ar).

## Integrantes

- [Cristian Gabriel Douce Suarez](https://github.com/cristiandouce)
- [Rafael Putaro](https://github.com/rafaelputaro)
- [Nahuel Callalli Rivera](https://github.com/calli97)

## De qué trata

Utilizando los Datasets del proyecto de [Buenos Aires Data](https://data.buenosaires.gob.ar/dataset/) construimos una REST API en Go para consultas de las siguientes 3 entidades del año 2021:

- [Usuarios](https://data.buenosaires.gob.ar/dataset/bicicletas-publicas/resource/c560b546-67fc-4017-a405-edbd36eec8f6)
- [Bicicleteros](https://data.buenosaires.gob.ar/dataset/estaciones-bicicletas-publicas/resource/e62fb076-dc64-4c42-b8ea-a8dd47395062)
- [Recorridos](https://data.buenosaires.gob.ar/dataset/bicicletas-publicas/resource/a9095876-e584-4b0d-976c-a4600455565b)

Implementamos los parsers utilizando los utilitarios nativos del lenguaje, e implementamos el servidor Web utilizando el framework [Gin](https://github.com/gin-gonic/gin).

## Cómo correrlo de forma local

Todo el código relativo al proyecto se encuentra en el directorio [./proyecto](./proyecto/)

Para correrlo de forma local solo es necesario correr

```sh
# Desde el root del respositorio
make run

# O bien, ingresando a la carpeta ./proyecto
go run main.go
```

## Cómo correr los tests

Se implementaron un set de tests para los parsers a modo de explorar los utilitarios que ofrece el lenguaje. Para correr los tests es necesario ejecutar:

```sh
# Desde el root del repositorio
make test

# O bien, ingresando a la carpeta ./proyecto
go test ./...
```

## Estructura del proyecto

- [./proyecto/api](./proyecto/api): Paquete de API que inicializa el WEB framework y define la rutas y los filtros.
- [./proyecto/config](./proyecto/config): Paquete de configuración que permite agregar configuracion desde variables de entorno.
- [./proyecto/model](./proyecto/model): Paquete de coleccion de las estructuras de los Modelos de los tipos de datos de los JSON y CSV parseados.
- [./proyecto/datasets](./proyecto/datasets): Locación de los datasets utilizados.
- [./proyecto/database](./proyecto/database): Realiza el parsing propiamente dicho de las fuentes de datos en CSV o JSON, y los expone para consumo interno como colección de tipos de datos.
- [./proyecto/logger](./proyecto/logger): Utilitario para logs del sistema.
- [./proyecto/utils](./proyecto/utils): Coleccion de utilitarios para el proyecto en general.

## Descripción de la API Rest

### [GET /api/bicicleteros](https://team-capudo-rest-api.herokuapp.com/api/bicicleteros)

Devuelve la lista total de bicicleteros de la Ciudad de Buenos Aires. Permite filtrar por query:

Parametro | Descripción | Ejemplo
--:|--|--
`anclajes_t` | Filtrado de bicicleterospor cantidad total y exacta de anclajes | `GET /api/bicicleteros?anclajes_t=30`
`anclajes_max` | Filtrado por cantidad máxima de anclajes | `GET /api/bicicleteros?anclajes_max=30`
`anclajes_min` | Filtrado por cantidad mínima de anclajes| `GET /api/bicicleteros?anclajes_min=30`

> Los parámetros de consulta pueden combinarse.

### [GET /api/bicicleteros/:id](https://team-capudo-rest-api.herokuapp.com/api/bicicleteros/2)

Devuelve la información de un bicicletero por `:id`.

Ejemplo:

```json
{
  "id": 2,
  "nombre_referencia": "2BAEcobici",
  "nombre_estacion": "002 - Retiro I",
  "ubicacion": "Ramos Mejia, Jose Maria, Dr. Av. & Del Libertador Av.",
  "tipo": "AUTOMÁTICA",
  "horario": "Estación automática: disponibilidad las 24 horas",
  "anclajes_t": 20,
  "latitud": -34.592422,
  "longitud": -58.37471
}
```

### [GET /api/recorridos](https://team-capudo-rest-api.herokuapp.com/api/recorridos)

Devuelve la lista total de recorridos de la Ciudad de Buenos Aires. Permite filtrar por query:

Parametro | Descripción | Ejemplo
--:|--|--
`id_estacion_origen` | Filtrado por estación de origen del recorrido | `GET /api/recorridos?id_estacion_origen=123BAEcobici`
`id_estacion_destino` | Filtrado por estación de destino del recorrido | `GET /api/recorridos?id_estacion_destino=123BAEcobici`
`fecha_desde` | Filtrado por fecha en que se inicio el recorrido | `GET /api/recorridos?fecha_desde=2020-02-20`
`fecha_hasta` | Filtrado por fecha en la que se finalizo el recorrido | `GET /api/recorridos?fecha_hasta=2020-02-20`
`id_usuario` | Filtrado por usuario que ha realizado el recorrido | `GET /api/recorridos?id_usuario=123BAEcobici`

> Los parámetros de consulta pueden combinarse.

### [GET /api/recorridos/:id](https://team-capudo-rest-api.herokuapp.com/api/recorridos/7737009BAEcobici)

Devuelve la información de un recorrido por `:id`.

Ejemplo:

```json
{
  "id": "7737009BAEcobici",
  "duracion_recorrido": 1658,
  "fecha_origen_recorrido": "2020-05-18T13:06:42Z",
  "id_estacion_origen": "33BAEcobici",
  "nombre_estacion_origen": "033 - Facultad de Medicina",
  "direccion_estacion_origen": "Pres. José Evaristo Uriburu 987",
  "longitud_estacion_origen": -58.39898,
  "latitud_estacion_origen": -34.59709,
  "fecha_destino_recorrido": "2020-05-18T13:34:20Z",
  "id_estacion_destino": "175BAEcobici",
  "nombre_estacion_destino": "147 - Constitución",
  "direccion_estacion_destino": "Avenida Juan de Garay 1050",
  "longitud_estacion_destino": -58.380707,
  "latitud_estacion_destino": -34.62685,
  "id_usuario": "211329BAEcobici",
  "modelo_bicicleta": "ICONIC"
}
```

### [GET /api/usuarios](https://team-capudo-rest-api.herokuapp.com/api/usuarios)

Devuelve la lista total de recorridos de la Ciudad de Buenos Aires. Permite filtrar por query:

Parametro | Descripción | Ejemplo
--:|--|--
`genero` | Filtrado por genero del usuario (`F\|M\|OTRO`) | `GET /api/usuarios?genero=M`
`edad` | Filtrado por edad exacta del usuario | `GET /api/usuarios?edad=37`
`edad_desde` | Filtrado por edad minima del usuario | `GET /api/usuarios?edad_desde=...`
`edad_hasta` | Filtrado por edad maxima del usuario | `GET /api/usuarios?edad_hasta=...`
`fecha_alta_desde` | Filtrado por fecha de alta minima | `GET /api/usuarios?fecha_alta_desde=2020-01-02`
`fecha_alta_hasta` | Filtrado por fecha de alta máxima | `GET /api/usuarios?fecha_alta_hasta=2021-12-31`

> Los parámetros de consulta pueden combinarse.

### [GET /api/usuarios/:id](https://team-capudo-rest-api.herokuapp.com/api/usuarios/605638)

Devuelve la información de un usuario por `:id`.

Ejemplo:

```json
{
  "id": 605638,
  "genero": "OTRO", // F | M | OTRO
  "edad": 27,
  "fecha_alta": "2020-09-29T13:03:00Z"
}
```

### [GET /api/usuarios/stats](https://team-capudo-rest-api.herokuapp.com/api/usuarios/stats)

Devuelve un objeto de estadísticas calculadas sobre la base de datos completa de usuarios.

Ejemplo:

```json
{
  "edad_maxima": 73,
  "edad_minima": 16,
  "edad_promedio": 32,
  "por_genero": {
    "hombre": {
      "edad_maxima": 73,
      "edad_minima": 17,
      "edad_promedio": 32
    },
    "mujer": {
      "edad_maxima": 71,
      "edad_minima": 16,
      "edad_promedio": 33
    },
    "otro": {
      "edad_maxima": 63,
      "edad_minima": 18,
      "edad_promedio": 31
    }
  }
}
```

### [GET /api/usuarios/stats_sync](https://team-capudo-rest-api.herokuapp.com/api/usuarios/stats_sync)

Identico al endpoint `GET /api/usuarios/stats` pero realiza los cálculos de forma sincrónica, en lugar de utilizar `goroutines` como lo hace el anterior.

## Heroku

El proyecto fue publicado en `dyno` gratuito de Heroku en la siguiente URL para poder experimentar y realizar pruebas de forma pública.

- [https://team-capudo-rest-api.herokuapp.com](https://team-capudo-rest-api.herokuapp.com)

> :info: Todos los endpoints mencionados arriba están disponibles
