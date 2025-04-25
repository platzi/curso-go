package main

import "fmt"

func modificarArreglo(arg *[5]int) {
	(*arg)[0] = 42
}

func main() {
	numeros := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numeros)

	modificarArreglo(&numeros)
	fmt.Println(numeros)

	numeros[0] = 10
	fmt.Println(numeros)
}