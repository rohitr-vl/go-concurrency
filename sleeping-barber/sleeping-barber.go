package sleepingbarber

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

// package level variables
// seating capacity of barber shop
var seatingCapacity = 7

// interval in milliseseconds in which next client arrives
var arrivalRate = 100

// time taken to do 1 haircut
var cutDuration = 1000 * time.Millisecond

// time for which the barbersop remains open and allows client to come
var timeOpen = 10 * time.Second

func SleepingBarber() {
	// seed random number generator for queueing clients/ working barbers
	// r := rand.New(rand.NewSource(time.Now().UnixMicro()))

	color.Yellow("--- Sleeping Barber Problem ---")

	// create channels
	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	// create barber shop
	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HairCutDuration: cutDuration,
		NumberOfBarbers: 0,
		ClientsChan:     clientChan,
		BarbersDoneChan: doneChan,
		Open:            true,
	}
	color.Green("The shop is open for the day!")
	// add barbers
	shop.addBarber("Frank")
	shop.addBarber("Chiku")
	shop.addBarber("Jigga")

	// start the barbershop as a goroutine
	// one channel to notify the barbershop to not accept any more clients, as it is closing now
	shopClosing := make(chan bool)
	// other channel to mark the shop as closed
	closed := make(chan bool)
	go func() {
		// keep the barber shop open for given time
		<-time.After(timeOpen)
		shopClosing <- true
		shop.closeShopForTheDay()
		closed <- true

	}()

	// add clients using go-routine
	i := 1
	go func() {
		for {
			// random frequency of adding client
			randMilliseconds := 50 + rand.Intn(300-50)

			select {
			case <-shopClosing:
				// receiving value from shopClosing channel means barbershop needs to be closed and
				// we can no longer add clients, so stop this go-routine
				return
			case <-time.After(time.Millisecond * time.Duration(randMilliseconds)):
				clientName := []string{"Client", strconv.Itoa(randMilliseconds), strconv.Itoa(i)}
				shop.addClient(strings.Join(clientName, "_"))
				i++
			}
		}
	}()

	// block until the barbershop is closed
	<-closed
	color.Cyan("Total Clients arrived:%d", totalClients)
	color.Cyan("Clients left without haircut:%d", clientsLeft)
	color.Cyan("Hair cut count by barber:")
	for name, cnt := range haircutCount {
		color.Cyan("%s : %d", name, cnt)
	}
}
