package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	fmt.Printf("Numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))

	neededNumbers := numbers[:len(numbers)-10]

	fmt.Printf("NeedNumbers = %v\n", neededNumbers)
	fmt.Printf("length = %d\n", len(neededNumbers))
	fmt.Printf("capacity = %d\n", cap(neededNumbers))

	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	fmt.Printf("NumbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))
}
