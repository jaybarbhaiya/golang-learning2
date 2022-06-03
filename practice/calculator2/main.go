package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	fmt.Println("----- Extended Calculator -----")

	floatValue1 := readValue("Value1:")
	floatValue2 := readValue("Value2:")

	fmt.Print("Choose operator (+, -, *, /):")
	operator, _ := reader.ReadString('\n')
	operator = strings.TrimSpace(operator)

	switch operator {
	case "+":
		add(floatValue1, floatValue2)
	case "-":
		subtract(floatValue1, floatValue2)
	case "*":
		multiply(floatValue1, floatValue2)
	case "/":
		divide(floatValue1, floatValue2)
	default:
		panic("Invalid Operator!")
	}
}

func readValue(inputKey string) float64 {
	fmt.Print(inputKey)
	value, _ := reader.ReadString('\n')
	floatValue, parseError := strconv.ParseFloat(strings.TrimSpace(value), 64)
	if parseError != nil {
		panic(parseError)
	}
	return floatValue
}

func add(value1, value2 float64) {
	result := value1 + value2
	fmt.Printf("%.2f + %.2f = %.2f\n", value1, value2, result)
}

func subtract(value1, value2 float64) {
	result := value1 - value2
	fmt.Printf("%.2f - %.2f = %.2f\n", value1, value2, result)
}

func multiply(value1, value2 float64) {
	result := value1 * value2
	fmt.Printf("%.2f * %.2f = %.2f\n", value1, value2, result)
}

func divide(value1, value2 float64) {
	result := value1 * value2
	fmt.Printf("%.2f / %.2f = %.2f\n", value1, value2, result)
}
