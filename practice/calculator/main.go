package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("-----Simple calculator-----")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter value1: ")
	value1, _ := reader.ReadString('\n')
	// convert value to float
	floatValue1, value1ParseErr := strconv.ParseFloat(strings.TrimSpace(value1), 64)
	if value1ParseErr != nil {
		panic(value1ParseErr)
	}

	fmt.Print("Enter value2: ")
	value2, _ := reader.ReadString('\n')
	// convert value to float
	floatValue2, value2ParseErr := strconv.ParseFloat(strings.TrimSpace(value2), 64)
	if value2ParseErr != nil {
		panic(value2ParseErr)
	}

	// add the float values
	floatSum := floatValue1 + floatValue2
	fmt.Printf("The sum of %v and %v is %v", floatValue1, floatValue2, floatSum)
}
