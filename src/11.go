package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Generate a random number between 1 and 10
	n := rand.Intn(10) + 1
	fmt.Println("Random number:", n)
}
