// converter/converter.go
package converter

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const apiURL = "https://api.exchangerate-api.com/v4/latest/"

// ExchangeRates holds the exchange rates for various currencies.
type ExchangeRates struct {
	Rates map[string]float64 `json:"rates"`
}

// GetExchangeRate fetches exchange rates for the given base currency.
func GetExchangeRate(baseCurrency string) (map[string]float64, error) {
	resp, err := http.Get(apiURL + baseCurrency)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch exchange rate: %s", resp.Status)
	}

	var rates ExchangeRates
	err = json.NewDecoder(resp.Body).Decode(&rates)
	if err != nil {
		return nil, err
	}

	return rates.Rates, nil
}

// ConvertCurrency converts the given amount from one currency to another using live exchange rates.
func ConvertCurrency(amount float64, fromCurrency string, toCurrency string) (float64, error) {
	rates, err := GetExchangeRate(fromCurrency)
	if err != nil {
		return 0, err
	}

	rate, exists := rates[toCurrency]
	if !exists {
		return 0, fmt.Errorf("conversion rate for %s not found", toCurrency)
	}

	return amount * rate, nil
}
