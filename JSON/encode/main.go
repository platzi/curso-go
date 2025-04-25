package main

import (
	"fmt"
	"encoding/json"
)

type Person struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email"`
}

func main() {
	person := Person{
        Name:  "Amin Espinoza",
        Age:   40,
        Email: "amin.espinoza@platzi.com",
    }

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	fmt.Println("Tu archivo JSON es:", string(jsonData))
}