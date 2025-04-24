package main

import "fmt"

func main() {

	var numero int = 10
	numero2 := 20
	fmt.Println(numero, numero2)

	var numeroEntero = 10
	var numeroDoble = 20.5
	resultado := numeroEntero + int(numeroDoble)
	fmt.Println(resultado)

	var nombre string = "Amin"
	apellido := "Gonzalez"
	nombreCompleto := nombre + " " + apellido
	fmt.Println(nombreCompleto)
}