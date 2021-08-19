package main

import (
	"fmt"
)

func main() {
	var number1, number2 int
	var op string
	var sum int = 0
	fmt.Scanf("%d\n", &number1)
	fmt.Scanf("%s\n", &op)
	fmt.Scanf("%d\n", &number2)

	switch op {
	case "+":
		sum = number1 + number2
	case "-":
		sum = number1 - number2
	case "*":
		sum = number1 * number2
	case "/":
		sum = number1 / number2
	default:
	}
	fmt.Printf("%d", sum)
}
