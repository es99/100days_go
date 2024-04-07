// O programa a seguir lê uma lista de compras do teclado
// e exibe item por item mostrando a quantidade de cada um
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	counts := make(map[string]int)

	fmt.Println("Digite os itens abaixo, ao final da lista, digite 'fim'.")
	for input.Scan() {
		if input.Text() == "fim" {
			break
		}
		counts[input.Text()]++
	}

	for k, v := range counts {
		fmt.Printf("%d\t%s\n", v, k)
	}
}
