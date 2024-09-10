// Calculate the factorial of a number using a recursive function in Go.
package main

import "fmt"

// factorial is a recursive function that calculates the factorial of a number.
func factorial(n int) int {
	if n < 0 {
		panic("and how can I get you the Gamma function?")
	}
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var symbols int
	var number int
	fmt.Print("Enter factorial number: ")
	symbols, err := fmt.Scanln(&number)
	if err != nil {
		fmt.Println("Invalid argument error:", err)
		return
	}
	result := factorial(number)
	fmt.Printf("Factorial of %d is %d\n symbols %d", number, result, symbols)
}
