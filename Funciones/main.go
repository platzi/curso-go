package main

import "fmt"

func suma(numero1 int, numero2 int) int {
	resultado := numero1 + numero2
	return resultado
}

func sumaLarga(numero1, numero2, numero3 int) int {
    return numero1 + numero2 + numero3
}

func main() { 
	var numero1, numero2, numero3 int

	fmt.Println("Ingresa el primer numero: ")
	fmt.Scanln(&numero1)
	fmt.Println("Ingresa el segundo numero: ")
	fmt.Scanln(&numero2)
	fmt.Println("Ingresa el tercer numero: ")
	fmt.Scanln(&numero3)

	resultado := suma(numero1, numero2)
	fmt.Println("La suma de los dos numeros es: ", resultado)

	resultadoLargo := sumaLarga(numero1, numero2, numero3)
	fmt.Println("La suma de los tres numeros es: ", resultadoLargo)
}