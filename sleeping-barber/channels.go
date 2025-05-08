package sleepingbarber

import (
	"fmt"
	"strings"
)

/*
In shout function parameters, to make PING a receive only channel, add arrow before chan keyword
and to make Pong a send only channel, add arrow after chan keyword
*/
func shout(ping <-chan string, pong chan<- string) {
	for {
		//receiving string message from channel ping
		str, ok := <-ping
		if !ok {
			fmt.Println("Error while receiving msg on channel Ping")
		}

		//sending string message to channel pong
		pong <- fmt.Sprintf("%s !!!", strings.ToUpper(str))
	}
}
func Channels() {
	ping := make(chan string)
	pong := make(chan string)
	go shout(ping, pong)

	fmt.Println("Type something and press enter, enter Q to quit:")
	for {
		fmt.Print("->")
		// get user input
		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if userInput == strings.ToLower("q") {
			break
		}
		ping <- userInput

		//wait for response
		response := <-pong
		fmt.Println("Response:", response)
	}
	fmt.Println("Closing channels")
	close(ping)
	close(pong)
}
