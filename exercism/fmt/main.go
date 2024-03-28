package main

import (
	"fmt"
	"partyrobot"
)

func main() {
	nome := "Engels"
	idade := 37
	companheira := "Alessandra"
	numeroMesa := 13
	direcao := "on the right"
	distancia := 15.78

	fmt.Println(partyrobot.Welcome(nome))
	fmt.Println(partyrobot.HappyBirthday(nome, idade))
	fmt.Println(partyrobot.AssignTable(nome, numeroMesa, companheira, direcao, distancia))
}
