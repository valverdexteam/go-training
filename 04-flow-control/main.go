package main

import (
	"fmt"
)

func main() {
	values := []string{"James", "John", "Paul", "Thiago"}

	if contains(values, "Thiago") {
		fmt.Println("I found Thiago on the array")
	} else {
		fmt.Println("Thiago was not found there")
	}

	for i := len(values) - 1; i >= 0; i-- {
		printPosition(values, values[i])
	}

}

func printPosition(array []string, content string) {
	indexOf := indexOf(array, content)
	switch indexOf {
	case 0:
		fmt.Printf("%s was found on first position\n", content)
		break
	case 1:
		fmt.Printf("%s was found on second position\n", content)
		break
	case 2:
		fmt.Printf("%s was found on third position\n", content)
		break
	case 3:
		fmt.Printf("%s was found on fourth position\n", content)
		break
	}
}

func contains(array []string, content string) bool {
	var indexOf = indexOf(array, content)
	return indexOf >= 0
}

func indexOf(array []string, content string) int {
	for i, element := range array {
		if element == content {
			return i
		}
	}
	return -1

}
