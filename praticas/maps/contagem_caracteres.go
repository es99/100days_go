/*
5 functions de nivel easy para praticar o conceito de maps
*/

package main

import (
	"fmt"
)

// Escreva uma função que receba um map como entrada, uma chave e um valor, e adicione o par chave-valor ao map.
func AdicionaElemento(dados map[string]float64, key string, value float64) {
	_, status := dados[key]
	if status {
		fmt.Printf("%s ja existe no dicionario\n", key)
	} else {
		dados[key] = value
		fmt.Printf("%s adicionado no dicionario com o valor %f\n", key, value)
	}
}

// ContaCaractere recebe uma string como entrada e retorna um map contendo a contagem de cada caractere na string.
func ContaCaractere(palavra string) map[rune]int {
	contagem := make(map[rune]int)
	for _, c := range palavra {
		contagem[c] += 1
	}
	return contagem
}

// Escreva uma função que receba um map como entrada e um valor(key) como parâmetro,
// e retorne verdadeiro se o valor estiver presente no map e falso caso contrário.
func ProcuraElemento(contatos map[string]int, valor string) bool {
	_, status := contatos[valor]
	if status {
		return true
	} else {
		return false
	}
}

func main() {
	// ContaCaractere()
	palavra := "abacate"

	// ProcuraElemento()
	agenda := map[string]int{
		"Engels":     37,
		"Alessandra": 46,
		"Kaoue":      32,
		"Francisco":  11,
	}

	// AdicionaElemento()
	prices := make(map[string]float64)

	AdicionaElemento(prices, "abacate", 3.50)
	AdicionaElemento(prices, "abacaxi", 5)
	AdicionaElemento(prices, "pera", 0.50)
	AdicionaElemento(prices, "abacaxi", 4.15)

	// ProcuraElemento()
	pessoa := "Francisco"

	// ProcuraElemento()
	if ProcuraElemento(agenda, pessoa) {
		fmt.Println("Encontrado na agenda.")
	} else {
		fmt.Println("Nao encontrado na agenda.")
	}

	// ContaCaractere()
	for k, v := range ContaCaractere(palavra) {
		fmt.Printf("%c - %d\n", k, v)
	}

	fmt.Println(prices)
}
