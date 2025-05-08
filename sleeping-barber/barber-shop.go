package sleepingbarber

import (
	"sync"
	"time"

	"github.com/fatih/color"
)

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientsChan     chan string
	Open            bool
}

// stat variables
var totalClients int
var clientsLeft int
var haircutCount = make(map[string]int)
var haircutCountLock sync.Mutex

func (shop *BarberShop) addBarber(barber string) {
	shop.NumberOfBarbers++

	go func() {
		isSleeping := false
		color.Yellow("%s goes to waiting room to check for clients", barber)
		haircutCount[barber] = 0
		// barber needs to keep checking for waiting clients or go to sleep, until the shop closes
		for {
			if len(shop.ClientsChan) == 0 {
				// no clients in the waiting room
				color.Yellow("%s goes to sleep", barber)
				isSleeping = true
			}
			// we do not use shop.Open because with multiple barbers it will cause a RACE problem
			// shopOpen is the 2nd parameter which is bool and indicates whether the value was sent to the channel and
			// the channel is still open. If this is false, channel is closed and empty
			client, shopOpen := <-shop.ClientsChan
			if shopOpen {
				if isSleeping {
					color.Yellow("Client %s wakes up Barber %s", client, barber)
					isSleeping = false
				}
				// cut hair
				shop.cutHair(barber, client)
			} else {
				// shop is closed, send barber home
				shop.sendBarberHome(barber)
				// closing the go-routine
				return
			}
		}
	}()
}

func (shop *BarberShop) cutHair(barber, client string) {
	color.Green("Barber %s is cutting %s hair", barber, client)
	time.Sleep(shop.HairCutDuration)
	haircutCountLock.Lock()
	haircutCount[barber]++
	haircutCountLock.Unlock()
	color.Green("Barber %s finished cutting %s hair", barber, client)
}

func (shop *BarberShop) sendBarberHome(barber string) {
	color.Cyan("barber %s is going home", barber)
	shop.BarbersDoneChan <- true
}

func (shop *BarberShop) closeShopForTheDay() {
	color.Cyan("Closing shop for the day")
	close(shop.ClientsChan)
	shop.Open = false

	// wait until all barbers are done
	for a := 1; a <= shop.NumberOfBarbers; a++ {
		<-shop.BarbersDoneChan
	}
	close(shop.BarbersDoneChan)
	color.Green("--- All barbers gone home, and barber shop closed ---")
}

func (shop *BarberShop) addClient(clientName string) {
	color.Green("%s arrives at barber shop", clientName)
	totalClients++
	if shop.Open {
		select {
		case shop.ClientsChan <- clientName:
			color.Yellow("%s takes a seat in waiting room", clientName)
		default:
			color.Red("The waiting room is full, so %s leaves", clientName)
			clientsLeft++
		}
	} else {
		color.Red("barber shop is closed, so %s leaves", clientName)
		clientsLeft++
	}
}
