package main

import (
	"fmt"
	"meuprojeto/blackjack"
)

func RetornaIdade(num int) string {
	switch {
	case num >= 0 && num < 1:
		return "recem-nascido"
	case num >= 1 && num < 12:
		return "crianca"
	case num >= 12 && num < 21:
		return "adolescente"
	case num >= 21 && num < 60:
		return "adulto"
	case num >= 60:
		return "idoso"
	default:
		return "opcao incorreta"
	}
}

func main() {
	carta1 := "ace"
	carta2 := "ace"
	dealerCard := "two"

	fmt.Println(blackjack.ParseCard(carta1))
	fmt.Println(blackjack.FirstTurn(carta1, carta2, dealerCard))
}
