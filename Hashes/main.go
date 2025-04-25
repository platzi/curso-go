package main

import (
	"fmt"
	"crypto/sha256"
)

func main() { 

	cadenaBase := "Amin Espinoza"
	
	hash := sha256.New()

	hash.Write([]byte(cadenaBase))

	arregloHash := hash.Sum(nil)
	fmt.Println(cadenaBase)
	fmt.Printf("%x\n", arregloHash)
}