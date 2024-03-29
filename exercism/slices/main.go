package main

import "fmt"

func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)

	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			return
		}
	}

	fmt.Println(num, "not found in", nums)
}

func main() {

	lista := []int{90, 91, 95}

	find(89, lista...)
	find(45, 56, 67, 45, 90, 109)
	find(87)

}
