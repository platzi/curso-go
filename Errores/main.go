package main

import (
	"fmt"
	"errors"
)

var ErrorDeCafe = fmt.Errorf("ya no hay cafe")
var ErrorDeEnergia = errors.New("ya no hay electricidad")

func hacerCafe(args int) error {
	if args == 2 {
		return ErrorDeCafe
	} else if args == 4 {
		return fmt.Errorf("haciendo cafe: %w", ErrorDeEnergia)
	}
	return nil
}


func main() {
	for i := range 5 {
        if err := hacerCafe(i); err != nil {

            if errors.Is(err, ErrorDeCafe) {
				fmt.Println(err)
                fmt.Println("Necesitamos más café")
            } else if errors.Is(err, ErrorDeEnergia) {
                fmt.Println("Ahora no puedo calentar agua")
            } else {
                fmt.Printf("Error desconocido: %s\n", err)
            }
            continue
        }

        fmt.Println("¡Ya hay café!")
    }
}