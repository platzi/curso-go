package main

import "fmt"

type ServerState int

const (
    StateIdle ServerState = iota
    StateConnected
    StateError
    StateRetrying
)

var stateName = map[ServerState]string{
    StateIdle:      "inactivo",
    StateConnected: "conectado",
    StateError:     "error",
    StateRetrying:  "retrying",
}

func (estadoServidor ServerState) String() string {
    return stateName[estadoServidor]
}

func main() {

	redServidor := verificacionDeRed(StateIdle)
	fmt.Println("Estado del servidor:", redServidor)

	segundRevision := verificacionDeRed(redServidor)
	fmt.Println("Estado del servidor:", segundRevision)

}

func verificacionDeRed(servidor ServerState) ServerState {
	switch servidor {
    case StateIdle:
        return StateConnected
    case StateConnected, StateRetrying:
        return StateIdle
    case StateError:
        return StateError
    default:
        panic(fmt.Errorf("estado desconocido: %s", servidor))
    }
}