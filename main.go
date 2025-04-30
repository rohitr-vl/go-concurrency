package main

import (
	"fmt"
	"sync"

	"go-concurrency/income"
	producerconsumer "go-concurrency/producer-consumer"
)

var msg3 string
var wg3 sync.WaitGroup

func printSomething(index int, value string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("%d: %s\n", index, value)
}

func updateMsg(s string, m *sync.Mutex) {
	defer wg3.Done()

	m.Lock()
	msg3 = s
	m.Unlock()
}

func main() {
	module := "prod-cons"
	switch module {
	case "prod-cons":
		producerconsumer.ProducerConsumer()
	case "income":
		income.IncomeCalc()
	case "wait_group":
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
	case "challenge_1":
		challenge()
	case "mutex_intro":
		msg3 = "Hello world"
		var mutex sync.Mutex
		wg3.Add(2)
		go updateMsg("Hello Universe", &mutex)
		go updateMsg("Hello Cosmos", &mutex)
		wg3.Wait()
		fmt.Println(msg3)
	default:
		fmt.Println("Please select a valid option!")
	}
}
