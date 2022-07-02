package usuarios

import (
	"capudo/model"
	"fmt"
)

type EdadStats map[string]int
type GeneroStats map[string]EdadStats

var GenerosMap = map[string]string{
	"mujer":  "F",
	"hombre": "M",
	"otro":   "OTRO"}

func CalcularEdadPromedio(usuarios *[]model.Usuario, ch chan int) {
	ch <- CalcularEdadPromedioSync(usuarios)
}

func CalcularEdadPromedioSync(usuarios *[]model.Usuario) int {
	suma := 0
	total := len(*usuarios)
	for _, usuario := range *usuarios {
		suma += int(usuario.GetEdadUsuario())
	}

	// Evitamos dividir por 0
	if total == 0 {
		total = 1
	}

	return suma / total
}

func CalcularEdadMaxima(usuarios *[]model.Usuario, ch chan int) {
	ch <- CalcularEdadMaximaSync(usuarios)
}

func CalcularEdadMaximaSync(usuarios *[]model.Usuario) int {
	max := -1
	for _, usuario := range *usuarios {
		if int(usuario.GetEdadUsuario()) > max {
			max = int(usuario.GetEdadUsuario())
		}
	}
	return max
}

func CalcularEdadMinima(usuarios *[]model.Usuario, ch chan int) {
	ch <- CalcularEdadMinimaSync(usuarios)
}

func CalcularEdadMinimaSync(usuarios *[]model.Usuario) int {
	min := 99999
	for _, usuario := range *usuarios {
		if int(usuario.GetEdadUsuario()) < min {
			min = int(usuario.GetEdadUsuario())
		}
	}
	return min
}

func CalcularEdadPorGenero(usuarios *[]model.Usuario, ch chan GeneroStats) {
	stats := make(GeneroStats)
	generos := []string{"mujer", "hombre", "otro"}

	for _, genero := range generos {
		ch1, ch2, ch3 := make(chan int), make(chan int), make(chan int)

		fmt.Println("Filtramos los usuarios por genero", genero)
		usuariosGenero := FiltrarPorGenero(usuarios, GenerosMap[genero])

		fmt.Println("Calculamos edad por genero", generos)
		go CalcularEdadPromedio(&usuariosGenero, ch1)
		go CalcularEdadMaxima(&usuariosGenero, ch2)
		go CalcularEdadMinima(&usuariosGenero, ch3)

		stats[genero] = EdadStats{
			"edad_promedio": <-ch1,
			"edad_maxima":   <-ch2,
			"edad_minima":   <-ch3}
	}

	fmt.Println("Terminamos el calculo de edad por genero", stats)

	// repetir lo mismo pero para arrays filtrados for genero y construir el mapa
	ch <- stats
}

func CalcularEdadPorGeneroSync(usuarios *[]model.Usuario) GeneroStats {
	stats := make(GeneroStats)
	generos := []string{"mujer", "hombre", "otro"}

	for _, genero := range generos {
		fmt.Println("Filtramos los usuarios por genero", genero)
		usuariosGenero := FiltrarPorGenero(usuarios, GenerosMap[genero])

		fmt.Println("Calculamos edad por genero", genero)
		stats[genero] = EdadStats{
			"edad_promedio": CalcularEdadPromedioSync(&usuariosGenero),
			"edad_maxima":   CalcularEdadMaximaSync(&usuariosGenero),
			"edad_minima":   CalcularEdadMinimaSync(&usuariosGenero)}
	}

	fmt.Println("Terminamos el calculo de edad por genero", stats)

	// repetir lo mismo pero para arrays filtrados for genero y construir el mapa
	return stats
}
