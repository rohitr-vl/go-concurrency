package producerconsumer

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NumberOfPizzas = 15

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NumberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("\n Received order# %d", pizzaNumber)

		rnd := rand.Intn(16) + 1
		msg := ""
		success := false

		// random logic to decide if order was sucess or fail
		if rnd < 7 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++
		fmt.Printf("\n Making pizza order# %d. It will take %d seconds.. \n", pizzaNumber, delay)

		//delay
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("\n*** FAILED: We ran out of ingredients for pizza order# %d!", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("\n*** FAILED: The cook quit while making pizza order# %d!", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("\n SUCCESS: Pizza order# %d is ready!", pizzaNumber)
		}
		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}
		return &p
	}
	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}
func pizzeria(pizzaMaker *Producer) {
	// keep track of which pizza we are making
	var i = 0

	// run forever or until we recive a quit notification
	for {
		// try to make pizzas
		currentPizza := makePizza(i)
		if currentPizza != nil {
			// due to concurrent go-routines, pizzaNumber can be any value from Consumer go-routine
			i = currentPizza.pizzaNumber
			select {
			case pizzaMaker.data <- *currentPizza:
				// we tried to make a pizza (so sent valid data to the data channel)
			case quitChan := <-pizzaMaker.quit:
				// close channels
				close(pizzaMaker.data)
				close(quitChan)
				// exit the go-routine
				return
			}
		}
	}
}

func ProducerConsumer() {
	//seed the random number generator
	rand.New(rand.NewSource(time.Now().UnixNano()))

	//print out message with color
	color.Cyan("The pizzeria is open for business")
	color.Cyan("--------------------------------")

	// create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	// run producer in background
	go pizzeria(pizzaJob)

	// create and run consumer
	for i := range pizzaJob.data {
		if i.pizzaNumber <= NumberOfPizzas {
			// if currently consumed order# is less than total required orders,
			// then orders are still pending to be processed
			if i.success {
				color.Green(i.message)
				color.Green("\n Order# %d is out for delivery!", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("Customer is mad about order# %d being cancelled!", i.pizzaNumber)
			}
		} else {
			// all pizza orders have been processed
			color.Cyan("Done making pizzas!!!")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("\n *** Error closing a channel", err)
			}
		}
	}
	// print out the ending message after both Consumer & Producer go-routines are done
	color.Cyan("------Done for the day------")
	color.Cyan("Made %d pizzas, Failed %d pizzas, Total pizzas: %d", pizzasMade, pizzasFailed, total)

	switch {
	case pizzasFailed > 10:
		color.Red("Terrible day!")
	case pizzasFailed > 7:
		color.Red("Bad day!")
	case pizzasFailed > 5:
		color.Red("OK day!")
	case pizzasFailed > 3:
		color.Red("Good day!")
	default:
		color.Red("Best day!")
	}
}
