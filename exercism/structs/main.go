package main

import (
	"fmt"
)

type Endereco struct {
	rua string
	cep string
}

type Pessoa struct {
	name        string
	age         int
	nationality string
	Endereco
}

func newPessoa(nome, nacionalidade, rua, cep string, idade int) Pessoa {
	return Pessoa{
		name:        nome,
		age:         idade,
		nationality: nacionalidade,
		Endereco:    Endereco{rua, cep},
	}
}

func main() {
	fulano := newPessoa("Engels", "brasileiro", "Av Epitacio", "58032-000", 37)
	sicrano := Pessoa{name: "Diego", nationality: "ucraniano", age: 40, Endereco: Endereco{"Cruz das armas", "67890-000"}}

	fmt.Println(fulano)
	fmt.Println(sicrano)
}
