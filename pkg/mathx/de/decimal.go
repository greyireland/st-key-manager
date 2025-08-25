package de

import (
	"github.com/shopspring/decimal"
)

func FromString(v string) decimal.Decimal {
	ret, _ := decimal.NewFromString(v)
	return ret
}
