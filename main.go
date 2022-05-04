package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/7kFour/go-dice-roller/pkg/randEngine"
)

func main() {

	// generate see for rand using current system time
	s1 := rand.NewSource(time.Now().UnixNano())

	for i := 0; i < 100; i++ {
		fmt.Println(randEngine.RandGen(s1, 6) + 1)
	}
}
