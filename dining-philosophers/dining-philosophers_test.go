package diningphilosophers

import (
	"testing"
	"time"
)

func Test_dine(t *testing.T) {
	eatTime = 0 * time.Second
	sleepTime = 0 * time.Second
	thinkTime = 0 * time.Second

	for i := 0; i < 10; i++ {
		// package level var, so no need to declare again
		finishTime = []string{}
		dine()
		if len(finishTime) != 5 {
			t.Errorf("Incorrect length of slice, 5!=%d", len(finishTime))
		}
	}
}
func Test_dineWithVaryingDelays(t *testing.T) {
	var theTests = []struct {
		name  string
		delay time.Duration
	}{
		{"quater second delay", time.Millisecond * 250},
		{"half second delay", time.Millisecond * 500},
		{"one second delay", time.Second},
	}

	for _, val := range theTests {
		finishTime = []string{}
		eatTime = val.delay
		sleepTime = val.delay
		thinkTime = val.delay
		dine()
		if len(finishTime) != 5 {
			t.Errorf("\n %s: Incorrect length of slice, 5!=%d", val.name, len(finishTime))
		}
	}
}
