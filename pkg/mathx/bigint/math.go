package bigint

import "math/big"

func DivDown(x, y *big.Int) *big.Int {
	return new(big.Int).Quo(x, y)
}

func DivUp(x, y *big.Int) *big.Int {
	return Add(ONE, DivDown(Sub(x, ONE), y))
}
