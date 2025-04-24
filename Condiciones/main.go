package main

import "fmt"

func main() {

	nombre := "Amin"
	edad := 40

	if nombre == "Amin" {
		fmt.Println("Hola Amin")
	} else {
		fmt.Println("Hola desconocido")
	}

	if edad >= 18 {
		fmt.Println("Ya puedes votar!")
	}

	if 9%2 == 0 {
		fmt.Println("El 9 es par")
	} else {
		fmt.Println("El 9 es impar")
	}

	if numero := 99; numero < 0 {
        fmt.Println(numero, "es negativo")
    } else if numero < 10 {
        fmt.Println(numero, "es un dígito")
    } else {
        fmt.Println(numero, "es un número mayor a 9")
    }
}