package currency

import (
	"github.com/shopspring/decimal"
	"testing"
)

func TestConvertCentsToDollars(t *testing.T) {
	cents := int64(1000)
	expectedDollars := decimal.NewFromFloat(10.00)

	actualDollars := ConvertCentsToDollars(cents)

	if !expectedDollars.Equal(actualDollars) {
		t.Errorf("Unexpected result. Expected: %s, got: %s", expectedDollars.String(), actualDollars.String())
	}
}

func TestConvertDollarsToCents(t *testing.T) {
	dollars := decimal.NewFromFloat(10.00)
	expectedCents := int64(1000)

	actualCents := ConvertDollarsToCents(dollars)

	if expectedCents != actualCents {
		t.Errorf("Unexpected result. Expected: %d, got: %d", expectedCents, actualCents)
	}
}
