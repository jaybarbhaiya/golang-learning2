package main

import (
	"fmt"
)

func main() {
	fmt.Println("Defining and calling functions")

	doSomething()

	sum := add2Values(2, 2)
	fmt.Println("Sum:", sum)

	multiSum, multiCount := addMultiValues(2, 2, 4, 2, 1)
	fmt.Println("Sum multi:", multiSum)
	fmt.Println("count multi:", multiCount)
}

func doSomething() {
	fmt.Println("Doing something")
}

func add2Values(value1, value2 int) int {
	return value1 + value2
}

func addMultiValues(values ...int) (int, int) {
	total := 0
	for _, v := range values {
		total += v
	}
	return total, len(values)
}
