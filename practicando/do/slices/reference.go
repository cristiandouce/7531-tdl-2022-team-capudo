package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("arr", arr)

	slc := arr[2:4]

	fmt.Println("slc", slc)

	// acá me modifica el valor del slice y del array
	slc[0] = 1

	fmt.Println("slc", slc)
	fmt.Println("arr", arr)
}
