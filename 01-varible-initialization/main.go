package main

import "fmt"

func main() {
	var firstVar = "first var"
	fmt.Printf("firstVar is of type %T and its value is [%s]\n", firstVar, firstVar)

	var firstIntegerA, firstIntegerB int = 5, 6
	fmt.Printf("firstIntegerA is of type %T and its value is [%d]\n", firstIntegerA, firstIntegerA)
	fmt.Printf("firstIntegerB is of type %T and its value is [%d]\n", firstIntegerB, firstIntegerB)

	var firstFloat float64 = 20.5
	fmt.Printf("firstFloat is of type %T and its value is [%f]\n", firstFloat, firstFloat)

	stringDeclaredWithShorthand := "String declared and initialiazed with shorthand"
	fmt.Printf("stringDeclaredWithShorthand is of type %T and its value is [%s]\n", stringDeclaredWithShorthand, stringDeclaredWithShorthand)

}
