package mathHelpers

import (
	"math/rand"
	"time"
)

func RandomInt(max int) int {
	return randomInt(0, max)
}

func randomInt(min, max int) int {
	return rand.Intn(max-min) + min
}

func RandomBetween(min, max int) int {
	rand.NewSource(int64(max))
	for i := range rand.Intn(100) {
		rand.NewSource(int64(i))
	}
	return rand.Intn(max-min) + min
}

func CoinToss() bool {
	rand.NewSource(time.Now().UnixNano())
	return rand.Intn(2) == 1
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
// The Min function returns the minimum value between two integers.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
