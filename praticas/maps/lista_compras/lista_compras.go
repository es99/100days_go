// Programa que lê uma lista de itens com suas respectivas quantidades e depois retorna o valor da conta de acordo
// com uma tabela de preços definida.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	lista := make(map[string]int)
	precos := map[string]float64{
		"laranja": 1.25,
		"abacate": 3.60,
		"tomate":  0.75,
		"cebola":  0.25,
		"alho":    0.05,
	}

	for input.Scan() {
		lista[input.Text()]++
	}

	for k, v := range lista {
		fmt.Printf("%s - quantidade: %d\n", k, v)
	}
	fmt.Printf("Valor total: %.2f\n", somaPrecos(lista, precos))

}

func somaPrecos(lista map[string]int, precos map[string]float64) float64 {
	var soma float64
	for fruta, qtd := range lista {
		for frutaListaPrecos, valor := range precos {
			if frutaListaPrecos == fruta {
				soma += float64(qtd) * valor
			}
		}
	}
	return soma
}
