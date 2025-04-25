package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {

	const saludo = "สวัสดี"

	fmt.Println("El saludo es:", saludo)
	fmt.Println("Len:", len(saludo))

	for i := 0; i < len(saludo); i++ {
        fmt.Printf("%x ", saludo[i])
    }

	fmt.Println("Conteo de runas:", utf8.RuneCountInString(saludo))

	for idx, valorRuna := range saludo {
        fmt.Printf("%#U comienza en %d\n", valorRuna, idx)
    }

}