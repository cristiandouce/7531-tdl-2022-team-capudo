package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5}

	slc := arr[2:4]

	// imprime:
	//   arr [1 2 3 4 5]
	//   slc [3 4]
	fmt.Println("arr", arr)
	fmt.Println("slc", slc)

	// acá me modifica el valor del slice y del array
	slc[0] = 1

	// imprime:
	//	slc [1 4]
	//	arr [1 2 1 4 5]
	fmt.Println("slc", slc)
	fmt.Println("arr", arr)

	// como se ve, modificó el valor del array de la posición 3, que es la que el slice referencia con la 0.
}
