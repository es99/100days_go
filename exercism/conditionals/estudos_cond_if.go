package main

import "fmt"

func ComparaVal(variavel string) string {
	var retorno string

	if variavel == "val" {
		retorno = "was val"
	}

	return retorno
}

func main() {
	num := 7

	if v := 1 * num; v > 10 {
		fmt.Println("v maior que 10:", v)
	} else {
		fmt.Println("v menor que 10, mostrando num:", num)
	}
}
