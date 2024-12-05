package main

import (
	"fmt"
	"github.com/ABHIJEET-MUNESHWAR/Currency-Exchange/internal/currency"
	"sync"
	"time"
)

func main() {
	ce := &currency.MyCurrencyExchange{
		Currencies: make(map[string]currency.Currency),
	}
	err := ce.FetchAllCurrencies()
	if err != nil {
		panic(err)
	}

	wg := sync.WaitGroup{}
	startTime := time.Now()

	for code := range ce.Currencies {
		wg.Add(1)
		go func(code string) {
			rates, err := currency.FetchCurrencyRates(code)
			if err != nil {
				panic(err)
			}
			ce.Currencies[code] = currency.Currency{
				Code:  code,
				Name:  ce.Currencies[code].Name,
				Rates: rates,
			}
			wg.Done()
		}(code)

	}
	wg.Wait()
	endTime := time.Now()
	fmt.Println("============== Results ==============")
	for _, curr := range ce.Currencies {
		fmt.Printf("%s (%s): %d rates\n", curr.Name, curr.Code, len(curr.Rates))
	}
	fmt.Println("=====================================")
	fmt.Println("Time taken: ", endTime.Sub(startTime))
}