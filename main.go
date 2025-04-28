package main

import (
	"fmt"
	"sync"
)

func printSomething(index int, value string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%d: %s\n", index, value)
}

/*
func main() {
	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"gamma",
		"delta",
		"sigma",
		"omega",
		"pi",
		"epsilon",
	}
	wg.Add(len(words))
	for i, val := range words {
		go printSomething(i, val, &wg)
	}
	wg.Wait()
	wg.Add(1)
	printSomething(10, "Last Line", &wg)
}
*/

func main() {
	challenge()
}
