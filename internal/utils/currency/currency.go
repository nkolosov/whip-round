package currency

import "math"

// ConvertCentsToDollars converts balance from cents to dollars
func ConvertCentsToDollars(cents int64) float64 {
	return float64(cents) / 100
}

// ConvertDollarsToCents converts balance from dollars to cents
func ConvertDollarsToCents(dollars float64) int64 {
	return int64(math.Round(dollars * 100))
}
