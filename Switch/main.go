package main

import (
    "fmt"
	"time"
)

func main() {

	i := 3
	fmt.Println("Escribe ", i, " como")

	switch i {
	case 1:
		fmt.Println("Uno")
	case 2:
		fmt.Println("Dos")
	case 3:
		fmt.Println("Tres")
	}

	switch time.Now().Weekday() {
		case time.Saturday, time.Sunday:
			fmt.Println("A descansar!!!")
		default:
			fmt.Println("Toca grabar más cursos en Platzi")
	}

	tiempo := time.Now()
	fmt.Println("La hora es", tiempo)

    switch {
    case tiempo.Hour() < 12:
        fmt.Println("Debes decir buenos días")
    default:
        fmt.Println("Debes decir buenas tardes")
    }
}