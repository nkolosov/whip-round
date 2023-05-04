package currency

import (
	"math"
	"testing"
)

func TestConvertCentsToDollars(t *testing.T) {
	cents := int64(10000)
	expectedDollars := float64(cents) / 100
	dollars := ConvertCentsToDollars(cents)

	if dollars != expectedDollars {
		t.Errorf("ConvertCentsToDollars(%v) returned %v, expected %v", cents, dollars, expectedDollars)
	}
}

func TestConvertDollarsToCents(t *testing.T) {
	dollars := 100.50
	expectedCents := int64(math.Round(dollars * 100))
	cents := ConvertDollarsToCents(dollars)

	if cents != expectedCents {
		t.Errorf("ConvertDollarsToCents(%v) returned %v, expected %v", dollars, cents, expectedCents)
	}
}
