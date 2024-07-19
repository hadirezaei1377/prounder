package calculator

import "fmt"

func main() {

	var operation string
	fmt.Print("Please select an operation: +, -, *, / : ")
	fmt.Scanln(&operation)

	var num1 int
	fmt.Print("Please input the first number: ")
	fmt.Scanln(&num1)

	var num2 int
	fmt.Print("Please input the second number: ")
	fmt.Scanln(&num2)

	switch operation {
	case "+":

		fmt.Print("Result: ", "\n", num1+num2)

	case "-":
		fmt.Print("Result: ", "\n", num1-num2)

	case "*":
		fmt.Print("Result: ", "\n", num1*num2)

	case "/":
		fmt.Print("Result: ", "\n", num1/num2)

	default:
		fmt.Println("Invalid operation selected. Please try again!")

	}
}
