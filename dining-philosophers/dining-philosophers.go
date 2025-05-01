package diningphilosophers

import (
	"sync"
	"time"

	"go-concurrency/helpers"

	"github.com/fatih/color"
)

type Philosopher struct {
	name      string
	leftFork  int
	rightFork int
}

var philosophers = []Philosopher{
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Aristotle", leftFork: 1, rightFork: 2},
	{name: "Pascal", leftFork: 2, rightFork: 3},
	{name: "Locke", leftFork: 3, rightFork: 4},
}

var hunger = 3
var eatTime = 1 * time.Second
var thinkTime = 0 * time.Second
var sleepTime = 0 * time.Second

func DiningPhilosophers() {
	// start with empty table and no one starts to eat until all are on the table
	color.Green("Dining Philosophers Problem")
	color.Green("---------------------------")

	// gather all and start the meal
	dine()

	// finished eating
	color.Green("The table is empty i.e. all go-routines closed")
}

func dine() {
	// wait group for each philosopher to eat
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	// waitgroup for all philosophers to get seated and ready to eat
	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	// map of 5 forks
	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	for i := 0; i < len(philosophers); i++ {
		// fire off a go-routine for current philosopher
		go diningProblem(philosophers[i], wg, forks, seated)
	}
	wg.Wait()
}

func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	//seat the philosopher at the table
	color.Blue("\n %s: %s is seated at the table", helpers.GetLocalTime(), philosopher.name)
	seated.Done()

	// eat three times
	for i := hunger; i > 0; i-- {
		/* logical race problem cannot be detected by "#go run -race ." command
		The logical race problem will occur when each philosopher has locked a single fork
		and are waiting for the other fork to be avaiable. this is in an infinite waiting
		and cannot be detected by -race option. In order to avoid this we make the philosopher
		choose the lower numbered fork first and so in case of Plato it is right fork first
		*/
		if philosopher.leftFork < philosopher.rightFork {
			forks[philosopher.leftFork].Lock()
			color.Yellow(" %s: %s takes the left fork #%d", helpers.GetLocalTime(), philosopher.name, philosopher.leftFork)
			forks[philosopher.rightFork].Lock()
			color.Yellow(" %s: %s takes the right fork #%d", helpers.GetLocalTime(), philosopher.name, philosopher.rightFork)
		} else {
			forks[philosopher.rightFork].Lock()
			color.Yellow(" %s: %s takes the right fork #%d", helpers.GetLocalTime(), philosopher.name, philosopher.rightFork)
			forks[philosopher.leftFork].Lock()
			color.Yellow(" %s: %s takes the left fork #%d", helpers.GetLocalTime(), philosopher.name, philosopher.leftFork)
		}

		color.Yellow(" %s: %s is eating with forks #%d & #%d", helpers.GetLocalTime(), philosopher.name, philosopher.leftFork, philosopher.rightFork)
		time.Sleep(eatTime)

		color.Yellow(" %s: %s is thinking with forks", helpers.GetLocalTime(), philosopher.name)
		time.Sleep(thinkTime)

		forks[philosopher.leftFork].Unlock()
		forks[philosopher.rightFork].Unlock()

		color.Yellow(" %s: %s has put down the forks #%d & #%d", helpers.GetLocalTime(), philosopher.name, philosopher.leftFork, philosopher.rightFork)
	}
	color.Blue(" %s: %s has eaten and left the table \n", helpers.GetLocalTime(), philosopher.name)
}
