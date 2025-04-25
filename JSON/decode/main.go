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
	
	jsonString :=`{"name":"Amin Espinoza","age":40,"email":"amin.espinoza@platzi.com"}`

	var person Person
	err := json.Unmarshal([]byte(jsonString), &person)
	if err != nil {
        fmt.Println("Error decoding JSON:", err)
        return
    }

	fmt.Println("Nombre:", person.Name)
	fmt.Println("Edad:", person.Age)
	fmt.Println("Email:", person.Email)
}