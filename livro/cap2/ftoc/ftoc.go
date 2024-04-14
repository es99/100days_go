// Ftoc cria uma funcao que recebe um float em fahrenheit e retorna em celsius.
package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%gF = %gC\n", freezingF, converteF(freezingF))
	fmt.Printf("%gF = %gC\n", boilingF, converteF(boilingF))
}

func converteF(num float64) float64 {
	return (num - 32) * 5 / 9
}
