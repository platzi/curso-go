package main

import (
    "fmt"
    "time"
)

func main() { 

	canal1 := make(chan string)
	canal2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		canal1 <- "mensaje uno"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		canal2 <- "mensaje dos"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msj1 := <-canal1:
			fmt.Println("Canal 1:", msj1)
		case msj2 := <-canal2:
			fmt.Println("Canal 2:", msj2)
		}
	}
}