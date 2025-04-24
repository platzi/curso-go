package main

import "fmt"

func main() {

	var arreglo[5]int
	fmt.Println("Arreglo completo:", arreglo)

	arreglo[4] = 100
	fmt.Println("Arreglo completo:", arreglo)
	fmt.Println("Elemento en la posición 4:", arreglo[4])

	fmt.Println("Tamaño del arreglo:", len(arreglo))

	lista := [5]int{1, 2, 3, 4, 5}
    fmt.Println("original:", lista)

	lista2 := [...]int{1, 2, 3, 4, 5, 6}
    fmt.Println("inferido:", lista2)
	fmt.Println("Tamaño del arreglo inferido:", len(lista2))
}