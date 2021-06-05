package util

import "math/rand"

func randomIntRange(min, max int) int {
	return rand.Intn(max-min) + min
}

func RandomNonZeroDigit() int {
	return randomIntRange(1, 9)
}
