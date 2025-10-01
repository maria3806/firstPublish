package main

import (
	"fmt"
)

func numbers() {
	var firstNum float64

	fmt.Print("Enter first number: ")
	fmt.Scanln(&firstNum)

	var secondNum float64
	fmt.Print("Enter second number: ")
	fmt.Scanln(&secondNum)

	var op string
	fmt.Print("Enter operation: (+, -, *, /) ")
	fmt.Scanln(&op)

	var result float64
	if op == "+" {
		result = firstNum + secondNum
	} else if op == "-" {
		result = firstNum - secondNum
	} else if op == "*" {
		result = firstNum * secondNum
	} else if op == "/" {
		if secondNum != 0 {
			result = firstNum / secondNum
		} else {
			fmt.Println("Error: cant divide to zero.")
			return
		}
	} else {
		fmt.Println("Sorry, i dont know what operation that is.")
		return
	}
	fmt.Printf("Result: %.2f\n", result)
}

func main() {
	numbers()
}
