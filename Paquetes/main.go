package main

import (
	"fmt"
	"os"
)

func main() {
	
	envVar := os.Getenv("HOME")
    if envVar == "" {
        fmt.Println("HOME environment variable is not set")
    } else {
        fmt.Printf("HOME environment variable: %s\n", envVar)
    }

	file, err := os.Create("ejemplo.txt")
    if err != nil {
        fmt.Printf("Error creating file: %v\n", err)
        return
    }
    defer file.Close()
}