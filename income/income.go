package income

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int64
}

func IncomeCalc() {
	var bankBalance int64
	var balance sync.Mutex
	fmt.Printf("\nThe initial bank balance is: %d.00", bankBalance)

	incomes := []Income{
		{Source: "main_job", Amount: 500},
		{Source: "gift_fam", Amount: 10},
		{Source: "part_time", Amount: 50},
		{Source: "investment", Amount: 100},
	}

	wg.Add(len(incomes))
	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				balance.Lock()
				tempBalance := bankBalance
				tempBalance += income.Amount
				bankBalance = tempBalance
				// fmt.Printf("\nIncome of week %d from %s is: $%d.00", week, income.Source, bankBalance)
				balance.Unlock()
			}
			// fmt.Printf("\n Annual Income from %s is: $%d.00", income.Source, bankBalance)
		}(i, income)
	}
	wg.Wait()
	fmt.Printf("\nFinal bank balance: $%d.00", bankBalance)
}
