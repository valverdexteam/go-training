package main

import (
	"fmt"
	"time"
)

func main() {

	result1 := printAlphabetWithChannel("First")

	result2 := printAlphabetWithChannel("Second")

	result := fanIn(result1, result2)

	for {
		msg := <-result
		if msg == "" {
			break
		}
		fmt.Println(msg)
	}

}

func printAlphabetWithChannel(processAlias string) <-chan string {
	result := make(chan string)
	go func() {
		i := 1
		for i <= 26 {
			time.Sleep(1 * time.Second)
			result <- fmt.Sprintf("%s : %q", processAlias, toChar(i))
			i++
		}
		close(result)
	}()

	return result
}

func toChar(i int) rune {
	return rune('A' - 1 + i)
}

func fanIn(result1, result2 <-chan string) <-chan string {
	result := make(chan string)
	go func() {
		for {
			select {
			case msg := <-result1:
				result <- msg
			case msg := <-result2:
				result <- msg
			}
		}
	}()

	return result
}
