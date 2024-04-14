// Aprendendo sobre retorno multiplos de funcoes, parametros e argumentos
package main

import (
	"fmt"
)

func main() {
	nota1 := 6.5
	nota2 := 7.6
	nota3 := 10.0
	var nome string
	nome = "Engels"

	fmt.Printf("A media do aluno %s: %.2f\n", nome, Media(nota1, nota2, nota3)) // passando 3 argumentos para a funcao Media, que recebe 3 parametros
	boasVindas, adeus := Greetings(nome)                                        // passando o argumento nome que será recebido pelo parametro name
	fmt.Println(boasVindas)
	fmt.Println(adeus)
}

// Media recebe 3 parametros e retorna 1 float (unico retorno)
func Media(x, y, z float64) float64 {
	return (x + y + z) / 3
}

// Greetings recebe 1 parametro e retorna 2 valores do tipo string (multiplo retorno)
func Greetings(name string) (string, string) {
	var welcome string = fmt.Sprintf("Hello %s, welcome!", name)
	var goodbye string = fmt.Sprintf("Goodbye %s, see you soon...", name)
	return welcome, goodbye
}
