package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {

	var bankbalance int

	fmt.Printf("Initial account balance : %d.00", bankbalance)
	fmt.Println()

	//define weekly revenue
	incomes := []Income{
		{Source: "Main Job", Amount: 20000},
		{Source: "Gifts", Amount: 3000},
		{Source: "Part-time Job", Amount: 10000},
		{Source: "Investment", Amount: 6000},
	}

	//looping through 52 week to check how much user made
	var m sync.Mutex
	wg.Add(len(incomes))

	for i, income := range incomes {

		go func(i int, income Income) {
			defer wg.Done()
			for week := 0; week < 52; week++ {
				m.Lock()
				temp := bankbalance
				temp += income.Amount
				bankbalance = temp
				fmt.Printf("On week %d,you earned %d.00 from %s\n", week, bankbalance, income.Source)
				m.Unlock()
			}
		}(i, income)
	}
	wg.Wait()

	fmt.Printf("----------------Final bankbalance : %d.00-----------", bankbalance)
	fmt.Println()

}
