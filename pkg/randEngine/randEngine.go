package randEngine

import "math/rand"

// RandGen generates and returns random number
func RandGen(n rand.Source, d int) int {
	r1 := rand.New(n)

	return r1.Intn(d)
}
