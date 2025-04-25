package main

import (
    "fmt"
    "time"
)

func function(from string) {

	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	function("modo directo")

	go function("modo goroutine")

	go func(mensaje string) {
        fmt.Println(mensaje)
    }("enviando un mensaje")

	time.Sleep(1 * time.Second)
	fmt.Println("fin del programa")
}