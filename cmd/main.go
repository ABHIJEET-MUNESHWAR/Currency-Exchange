package main

import (
	"fmt"
	"github.com/ABHIJEET-MUNESHWAR/Currency-Exchange/internal/currency"
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
	startTime := time.Now()

	for code := range ce.Currencies {
		rates, err := currency.FetchCurrencyRates(code)
		if err != nil {
			panic(err)
		}
		ce.Currencies[code] = currency.Currency{
			Code:  code,
			Name:  ce.Currencies[code].Name,
			Rates: rates,
		}
	}
	endTime := time.Now()
	fmt.Println("============== Results ==============")
	for _, curr := range ce.Currencies {
		fmt.Printf("%s (%s): %d rates\n", curr.Name, curr.Code, len(curr.Rates))
	}
	fmt.Println("=====================================")
	fmt.Println("Time taken: ", endTime.Sub(startTime))
}