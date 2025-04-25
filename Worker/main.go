package main

import (
    "fmt"
    "time"
)

func worker(id int, tareas <-chan int, resultados chan int) {
	for i := range tareas {
		fmt.Printf("worker %d: procesando tarea %d\n", id, i)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "tarea iniciada", i)
		resultados <- i * 2
	}
}

func main() {
	const numeroTareas = 5

	tareas := make(chan int, numeroTareas)
	resultados := make(chan int, numeroTareas)

	for w := 1; w <= 3; w++ {
        go worker(w, tareas, resultados)
    }

	for j := 1; j <= numeroTareas; j++ {
        tareas <- j
    }
    close(tareas)

	for a := 1; a <= numeroTareas; a++ {
        <-resultados
    }
}