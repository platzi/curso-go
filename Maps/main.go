package main

import (
	"fmt"
	"maps"
)

func main() {

	mapa := make(map[string]int)

	mapa["Amin"] = 4
	mapa["Espinoza"] = 8

	fmt.Println("mapa: ", mapa)

	version1 := mapa["Espinoza"]
	fmt.Println("version1: ", version1)

	version2 := mapa["Amin"]
	fmt.Println("version2: ", version2)

	fmt.Println("longitud del mapa:", len(mapa))

	_, dato := mapa["Amin"]
	fmt.Println("dato: ", dato)

	delete(mapa, "Amin")
	fmt.Println("mapa: ", mapa)

	clear(mapa)
    fmt.Println("mapa:", mapa)

	nuevoMapa1 := map[string]int{"Miranda": 1, "Espinoza": 2}
    fmt.Println("mapa1:", nuevoMapa1)
	
    nuevoMapa2 := map[string]int{"Miranda": 1, "Espinoza": 2}
	fmt.Println("mapa2:", nuevoMapa2)

	if maps.Equal(nuevoMapa1, nuevoMapa2) {
		fmt.Println("Los mapas son iguales")
	} 
}