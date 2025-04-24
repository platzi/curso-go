package main

import "fmt"

func valoresMultiples(nombre1, nombre2, apellido string) (map[string]string, map[string]string) {
	mapa1 := make(map[string]string)
	mapa2 := make(map[string]string)

	mapa1[nombre1] = apellido
	mapa2[nombre2] = apellido

	return mapa1, mapa2
}


func main() { 
	nombre1 := "Amin"
	nombre2 := "Miranda"
	apellido := "Espinoza"

	mapaPrimerNombre, mapaSegundoNombre := valoresMultiples(nombre1, nombre2, apellido)

	fmt.Println("Mapa con el primer nombre:", mapaPrimerNombre)
	fmt.Println("Mapa con el segundo nombre:", mapaSegundoNombre)

	_, nombre := valoresMultiples(nombre1, nombre2, apellido)
    fmt.Println(nombre)
}