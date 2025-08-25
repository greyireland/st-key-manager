package bigint

import "math/big"

func IsZero(x *big.Int) bool {
	return x.Cmp(ZERO) == 0
}

func IsPositive(x *big.Int) bool {
	return x.Cmp(ZERO) > 0
}

func IsNegative(x *big.Int) bool {
	return x.Cmp(ZERO) < 0
}

func IsEqual(x, y *big.Int) bool {
	return x.Cmp(y) == 0
}

func IsNotEqual(x, y *big.Int) bool {
	return x.Cmp(y) != 0
}

func IsLessThan(x, y *big.Int) bool {
	return x.Cmp(y) < 0
}

func IsLessThanOrEqual(x, y *big.Int) bool {
	return x.Cmp(y) <= 0
}

func IsGreaterThan(x, y *big.Int) bool {
	return x.Cmp(y) > 0
}

func IsGreaterThanOrEqual(x, y *big.Int) bool {
	return x.Cmp(y) >= 0
}
