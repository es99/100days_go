package main

import (
	"fmt"
)

func main() {
	food := "taco"
	number := 4.3242

	compactNumber := fmt.Sprintf("%.2f", number)

	fmt.Printf("Bring me a %s\n", food)
	fmt.Println(compactNumber)

}
