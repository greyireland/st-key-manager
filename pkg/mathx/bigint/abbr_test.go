package bigint

import (
	"fmt"
	"math/big"
	"testing"
)

func TestMax(t *testing.T) {
	fmt.Println(Max(big.NewInt(1), nil))
	fmt.Println(Max(nil, big.NewInt(1)))
	fmt.Println(Max(nil, nil))
	fmt.Println(Max(big.NewInt(1), big.NewInt(2)))
}
