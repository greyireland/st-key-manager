package de

import (
	"errors"
	"github.com/shopspring/decimal"
)

var DivisionPrecision = 16

func Sqrt(d decimal.Decimal) decimal.Decimal {
	s, _, _ := SqrtRound(d, int32(DivisionPrecision))
	return s
}

var ErrImaginaryResult = errors.New("the result of this operation is imaginary")

const SqrtMaxIter = 1000000

func SqrtRound(d decimal.Decimal, precision int32) (decimal.Decimal, bool, error) {
	var (
		maxError = decimal.New(1, -precision)
		one      = decimal.NewFromFloat(1)
		lo, hi   decimal.Decimal
	)

	// Handle cases where d < 0, d = 0, 0 < d < 1, and d > 1
	if d.GreaterThanOrEqual(one) {
		lo = decimal.Zero
		hi = d
	} else if d.Equal(one) {
		return one, true, nil
	} else if d.LessThan(decimal.Zero) {
		return decimal.Zero, false, ErrImaginaryResult
	} else if d.Equal(decimal.Zero) {
		return decimal.Zero, true, nil
	} else {
		// d is between 0 and 1. Therefore, 0 < d < Sqrt(d) < 1.
		lo = d
		hi = one
	}

	var mid decimal.Decimal
	for i := 0; i < SqrtMaxIter; i++ {
		mid = lo.Add(hi).Div(decimal.New(2, 0)) //mid = (lo+hi)/2;
		if mid.Mul(mid).Sub(d).Abs().LessThan(maxError) {
			return mid, true, nil
		}
		if mid.Mul(mid).GreaterThan(d) {
			hi = mid
		} else {
			lo = mid
		}
	}
	return mid, false, nil
}
