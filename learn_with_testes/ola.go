package main

import (
	"fmt"
)

// Ola retorna uma string
func Ola(nome string) string {
	return "Hello, " + nome
}

func Media(x, y, z float64) float64 {
	return (x + y + z) / 3
}

func main() {
	fmt.Println(Ola("Engels"))
	fmt.Println(Media(10, 75, 80))
}
