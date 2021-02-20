package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go printAlphabet(wg, "First")

	wg.Add(1)
	go printAlphabet(wg, "Second")

	wg.Wait()

}

func printAlphabet(waitGroup *sync.WaitGroup, processAlias string) {
	defer waitGroup.Done()
	i := 1
	for i <= 26 {
		time.Sleep(1 * time.Second)
		fmt.Println(fmt.Sprintf("%s : %q", processAlias, toChar(i)))
		i++
	}
}

func toChar(i int) rune {
	return rune('A' - 1 + i)
}
