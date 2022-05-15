package main

import (
	"fmt"
)

func main() {
	var action string
	var number1 int
	var number2 int
	fmt.Println("Calculator in GO language")
	fmt.Println("Write two numbers:")
	fmt.Scan(&number1, &number2)
	var result int
	for i := 1; i < 10; i++ {
		fmt.Println("Select an action(\"+\",\"-\",\"*\",\"/\",\"=\"):")
		fmt.Scan(&action)
		switch action {
		case "+":
			if result == 0 {
				result = Add(number1, number2)
				fmt.Println("Result: ", result)
			} else {
				fmt.Println("Write the second number")
				fmt.Scan(&number2)
				result = Add(result, number2)
				fmt.Println("Result: ", result)
			}
		case "-":
			if result == 0 {
				result := Sub(number1, number2)
				fmt.Println("Result: ", result)
			} else {
				fmt.Println("Write the second number")
				fmt.Scan(&number2)
				result = Sub(result, number2)
				fmt.Println("Result: ", result)
			}
		case "*":
			if result == 0 {
				result := Multiplication(number1, number2)
				fmt.Println("Result: ", result)
			} else {
				fmt.Println("Write the second number")
				fmt.Scan(&number2)
				result = Multiplication(result, number2)
				fmt.Println("Result: ", result)
			}
		case "/":
			if result == 0 {
				result := Division(number1, number2)
				fmt.Println("Result: ", result)
			} else {
				fmt.Println("Write the second number")
				fmt.Scan(&number2)
				result = Division(result, number2)
				fmt.Println("Result: ", result)
			}
		}
		if action == "=" {
			break
		}
	}
}
func Add(number1, number2 int) int {
	return number1 + number2
}
func Sub(number1, number2 int) int {
	return number1 - number2
}
func Multiplication(number1, number2 int) int {
	return number1 * number2
}
func Division(number1, number2 int) int {
	return number1 / number2
}
