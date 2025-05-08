package sleepingbarber

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	for {
		time.Sleep(6 * time.Second)
		ch <- "this is from server 1"
	}
}

func server2(ch chan string) {
	for {
		time.Sleep(3 * time.Second)
		ch <- "this is from server 2"
	}
}

func ChannelSelect() {
	fmt.Println("--Select with Channels--")

	channel1 := make(chan string)
	channel2 := make(chan string)

	go server1(channel1)
	go server2(channel2)

	for i := 0; i <= 9; i++ {
		select {
		case s1 := <-channel1:
			fmt.Println("Case One:", s1)
		case s2 := <-channel1:
			fmt.Println("Case Two:", s2)
		case s3 := <-channel2:
			fmt.Println("Case Three:", s3)
		case s4 := <-channel2:
			fmt.Println("Case Four:", s4)
			/*		default:
					// can be used to aviod deadlock
					break
			*/
		}
	}
}
