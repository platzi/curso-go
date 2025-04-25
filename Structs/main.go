package main

import "fmt"

type persona struct {
	nombre string
	edad   int
}

func nuevaPersona(nombre string, edad int) *persona {
	nuevoIndividuo := persona{ nombre: nombre}
	nuevoIndividuo.edad = 42
	return &nuevoIndividuo
}

func main() {
	fmt.Println(persona{"Marcela", 42 })

	fmt.Println(persona{nombre: "Santiago", edad: 44})

	fmt.Println(persona{nombre: "Felipe"})

	fmt.Println(nuevaPersona("Roberto", 50))

	personita := persona{nombre: "Miranda", edad: 7}
    fmt.Println(personita.nombre)

	edadPersonita := &personita
	fmt.Println(edadPersonita.edad)

	edadPersonita.edad = 8
    fmt.Println(edadPersonita.edad)
	fmt.Println(personita)
}