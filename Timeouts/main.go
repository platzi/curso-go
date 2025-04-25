package main

import (
    "fmt"
    "time"
)

func main() {

	contador1 := make(chan string, 1)

    go func() {
        time.Sleep(2 * time.Second)
        contador1 <- "resultado 1"
    }()

	select {
	case resultado := <-contador1:
		fmt.Println("Recibido:", resultado)
	case <- time.After(1 * time.Second):
		fmt.Println("timeout")
	}

	contador2 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        contador2 <- "resultado 2"
    }()

	select {
	case resultado := <-contador1:
		fmt.Println("Recibido:", resultado)
	case <- time.After(3 * time.Second):
		fmt.Println("timeout")
	}
}