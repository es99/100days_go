package main

import (
	"fmt"
)

const espanhol = "espanhol"
const prefixHelloEspanhol = "Hola, "
const prefixHelloPortuguese = "Hello, "

// Ola retorna uma string
func Ola(nome, idioma string) string {
	if nome == "" {
		nome = "world!"
	}
	if idioma == espanhol {
		return prefixHelloEspanhol + nome
	}
	return prefixHelloPortuguese + nome
}

func main() {
	fmt.Println(Ola("Engels", ""))
}
