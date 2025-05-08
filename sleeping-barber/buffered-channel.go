package sleepingbarber

import (
	"fmt"
	"time"
)

func listenToChannel(ch chan int) {
	i := <-ch
	fmt.Println("Got", i, "from channel")

	time.Sleep(1 * time.Second)
}

func BufferedChannels() {
	ch := make(chan int, 10)
	go listenToChannel(ch)
	for i := 0; i <= 100; i++ {
		fmt.Println("Sending", i, "to channel...")
		ch <- i
		fmt.Println("sent", i, "to channel")
	}
	fmt.Println("Done")
	close(ch)
}
