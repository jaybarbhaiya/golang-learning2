package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch")

	rand.Seed(time.Now().Unix())
	// dow := rand.Intn(7) + 1
	// fmt.Println("Day", dow)

	var result string
	switch dow := rand.Intn(7) + 1; dow {
	case 1:
		result = "Its Sunday"
		// fallthrough
	case 2:
		result = "Its Monday"
	default:
		result = "Some other day"
	}

	fmt.Println(result)
}
