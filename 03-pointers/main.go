package main

import (
	"fmt"
)

func main() {
	name := "Thiago"
	lastName := "Souza"
	completeName := fmt.Sprintf("%s %s", name, lastName)
	pointer := &completeName

	fmt.Println(fmt.Sprintf("This will print memory address %v", pointer))
	fmt.Println(fmt.Sprintf("This will print value %s", *pointer))

	printAll := func() {
		fmt.Println(fmt.Sprintf("name: %s | last name: %s | complete name: %s | pointer: %v | pointer value: %s", name, lastName, completeName, pointer, *pointer))
	}
	printAll()
	completeName = "Thiago Valverde de Souza"
	printAll()
}
