
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var mySlice []int
	for i := 0; i < 10; i++ {
		mySlice = append(mySlice, rand.Intn(10))
	}
	fmt.Println(mySlice)
}