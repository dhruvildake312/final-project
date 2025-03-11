package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// generate a random number between 1 and 10
	randomNumber := rand.Intn(10) + 1

	// print the result
	fmt.Println("Random number:", randomNumber)
}
