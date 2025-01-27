package support

import (
	r "math/rand"
	rand "math/rand/v2"
	"time"
)

func RandomInt(max int) int {
	return randomInt(0, max)
}

func randomInt(min, max int) int {
	return rand.IntN(max-min) + min
}

func RandomBetween(min, max int) int {
	r.NewSource(int64(max))
	for i := 0; i < r.Intn(100); i++ {
		r.NewSource(int64(i))
	}
	return r.Intn(max-min) + min
}

func CoinToss() bool {
	r.NewSource(time.Now().UnixNano())
	return r.Intn(2) == 1
}
