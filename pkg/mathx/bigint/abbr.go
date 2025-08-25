package bigint

import "math/big"

func Add(a, b *big.Int) *big.Int {
	return (&big.Int{}).Add(a, b)
}
func Sub(a, b *big.Int) *big.Int {
	return (&big.Int{}).Sub(a, b)
}

func Mul(a, b *big.Int) *big.Int {
	return (&big.Int{}).Mul(a, b)
}
func Div(a, b *big.Int) *big.Int {
	return (&big.Int{}).Div(a, b)
}
func Mod(a, b *big.Int) *big.Int {
	return (&big.Int{}).Mod(a, b)
}
func Sqrt(a *big.Int) *big.Int {
	return a.Sqrt(a)
}
func Equal(a, b *big.Int) bool {
	return a.Cmp(b) == 0
}
func Gt(a, b *big.Int) bool {
	return a.Cmp(b) > 0
}
func Lt(a, b *big.Int) bool {
	return a.Cmp(b) < 0
}
func Max(a, b *big.Int) *big.Int {
	if a == nil && b == nil {
		return ZERO
	}
	if a == nil && b != nil {
		return b
	}
	if a != nil && b == nil {
		return a
	}
	if Gt(a, b) {
		return a
	}
	return b
}
func Min(a, b *big.Int) *big.Int {
	if Lt(a, b) {
		return a
	}
	return b
}
func Pow(a, e *big.Int) *big.Int {
	return a.Exp(a, e, nil)
}
