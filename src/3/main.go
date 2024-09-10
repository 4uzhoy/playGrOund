// Find min and max number in a list of numbers
package main

import (
	"fmt"
)

// return the minimum and maximum number in a list of numbers
func findMinMax(numbers []int) (int, int, error) {
	if len(numbers) == 0 {
		return 0, 0, fmt.Errorf("empty list")
	}

	var min, max = numbers[0], numbers[0]

	for _, num := range numbers[1:] {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	return min, max, nil

}

func main() {
	numbers := []int{10, 15, 2, 30, 1, -5, 100, 0}
	// numbers := []int{}
	min, max, err := findMinMax(numbers)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Min: %d, Max: %d\n", min, max)
	}
}
