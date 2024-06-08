// converter/converter_test.go
package converter

import (
	"testing"
)

func TestConvertCurrency(t *testing.T) {
	amount := 100.0
	fromCurrency := "USD"
	toCurrency := "INR"

	convertedAmount, err := ConvertCurrency(amount, fromCurrency, toCurrency)
	if err != nil {
		t.Fatalf("Error converting currency: %v", err)
	}

	if convertedAmount <= 0 {
		t.Fatalf("Invalid converted amount: %v", convertedAmount)
	}

	t.Logf("Converted %f %s to %f %s", amount, fromCurrency, convertedAmount, toCurrency)
}
