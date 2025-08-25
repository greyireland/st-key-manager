package bigint

import (
	"fmt"
	"testing"
)

func TestFromString(t *testing.T) {
	a := FromString("0x123")
	fmt.Println(a)
}
