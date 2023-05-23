package currency

import (
	"github.com/shopspring/decimal"
)

// ConvertCentsToDollars converts balance from cents to dollars
func ConvertCentsToDollars(cents int64) decimal.Decimal {
	centsDecimal := decimal.NewFromInt(cents)
	dollars := centsDecimal.Div(decimal.NewFromInt(100))
	return dollars
}

// ConvertDollarsToCents converts balance from dollars to cents
func ConvertDollarsToCents(dollars decimal.Decimal) int64 {
	centsDecimal := dollars.Mul(decimal.NewFromInt(100))
	cents := centsDecimal.IntPart()
	return cents
}
