package main

import "fmt"

func main() {
	mensajes := make(chan string)

	go func() { mensajes <- "ping" }()

	mensaje := <-mensajes
	fmt.Println(mensaje)
	// El canal se puede cerrar

	go func() { mensajes <- "pong" }()

	segundoMensaje := <-mensajes
	fmt.Println(segundoMensaje)
	// El canal se puede cerrar
}