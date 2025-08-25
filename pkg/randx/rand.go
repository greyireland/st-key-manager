package randx

import (
	"math/rand"
	"time"
)

var (
	randSource = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func RandInt(n int64) int64 {
	return randSource.Int63n(n)
}
