package main

import (
	"fmt"
	"math"
)

func calc(operation string, num1 float64, num2 float64) {

	switch operation {
	case "add":
		result := num1 + num2
		fmt.Println("your result is:", result)

	case "sub":
		result := num1 - num2
		fmt.Println("your result is:", result)

	case "mul":
		result := num1 * num2
		fmt.Println("your result is:", result)

	case "div":
		result := num1 / num2

		if num2 == 0 {
			fmt.Println("Cannot divide by zero")
			break
		}
		fmt.Println("your result is:", result)

	case "mod":
		result := int(num1) % int(num2)
		fmt.Println("your result is:", result)

	case "pow":
		result := math.Pow(num1, num2)
		fmt.Println("your result is:", result)

	case "exit":
		fmt.Println("====Thank you for using our program====")

	default:
		fmt.Println("invalid command")

	}
}
