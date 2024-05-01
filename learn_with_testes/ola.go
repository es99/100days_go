package main

import (
	"fmt"
)

// Ola retorna uma string
func Ola(nome string) string {
	return "Hello, " + nome
}

func main() {
	fmt.Println(Ola("Engels"))
}
