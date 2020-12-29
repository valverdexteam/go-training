package main

import (
	"fmt"
)

func main() {
	a := 5
	b := 65
	result := Sum(a, b)
	fmt.Println(fmt.Sprintf("The sum of %d and %d is equal to %d", a, b, result))
}

func Sum(a, b int) int {
	return a + b
}
