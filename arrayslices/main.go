package main

import "fmt"

//Function to summall integers with an array or slice
func SumAll(numbers ...[]int) (result []int) {
	for _, num := range numbers {
		result = append(result, Sum(num))
	}
	return result
}

// Function to sum two integers
func Sum(numbers []int) int {
	var result int
	if len(numbers) > 0 {
		for i := range numbers {
			result += numbers[i]
		}
	} else {
		result = 0
	}
	return result
}
func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	zeros := []int{1, 1, 1, 1, 1, 1}
	fmt.Println(SumAll(numbers, zeros))
}
