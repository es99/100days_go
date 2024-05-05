package main

import (
	"fmt"
)

const japones = "japones"
const espanhol = "espanhol"
const frances = "frances"
const prefixHelloEspanhol = "Hola, "
const prefixHelloPortuguese = "Hello, "
const prefixHelloFrances = "Bonjour, "
const prefixHelloJapones = "Conichua, "

// Ola retorna uma string
func Ola(nome, idioma string) string {
	if nome == "" {
		nome = "world!"
	}

	return prefixoDeSaudacao(idioma) + nome
}

func prefixoDeSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case frances:
		prefixo = prefixHelloFrances
	case espanhol:
		prefixo = prefixHelloEspanhol
	case japones:
		prefixo = prefixHelloJapones
	default:
		prefixo = prefixHelloPortuguese
	}
	return
}

func main() {
	fmt.Println(Ola("Engels", "japones"))
}
