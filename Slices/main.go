package main

import (
    "fmt"
	"slices"
)

func main() {
	var arregloCadenas []string
    fmt.Println("datos:", "contenido", arregloCadenas, "condición:", arregloCadenas == nil, "longitud:", len(arregloCadenas) == 0)

	arregloCadenas = make([]string, 3)
	arregloCadenas[0] = "a"
    arregloCadenas[1] = "b"
    arregloCadenas[2] = "c"
	fmt.Println("datos:", arregloCadenas)
    fmt.Println("dato específico:", arregloCadenas[2])
    fmt.Println("longitud:", len(arregloCadenas))

	arregloCadenas = append(arregloCadenas, "d")
	arregloCadenas = append(arregloCadenas, "e", "f")
	fmt.Println("valores finales:", arregloCadenas)
	fmt.Println("longitud:", len(arregloCadenas))

	segundoArreglo := []string{"g", "h", "i"}
    fmt.Println("datos 2:", segundoArreglo)

	tercerArreglo := []string{"g", "h", "i"}
	
    if slices.Equal(segundoArreglo, tercerArreglo) {
        fmt.Println("2 es igual que 3")
    }
}