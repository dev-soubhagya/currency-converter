package main

import (
	"fmt"
	"log"

	"github.com/dev-soubhagya/currency-converter/converter"
)

func main() {
	amount := 100.0
	fromCurrency := "USD"
	toCurrency := "INR"

	convertedAmount, err := converter.ConvertCurrency(amount, fromCurrency, toCurrency)
	if err != nil {
		log.Fatalf("Error converting currency: %v", err)
	}

	fmt.Printf("%.2f %s is %.2f %s\n", amount, fromCurrency, convertedAmount, toCurrency)
}
