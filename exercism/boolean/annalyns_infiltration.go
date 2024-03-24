/*
Package main que inclui ferramentas através de funções que emulam um jogo de RPG onde

	a protagonista Annalyn tenta salvar sua melhor amiga das garras de sequestradores
	(arqueiro, cavaleiro).
*/
package main

import "fmt"

// CanFastAttack recebe um valor boleano para saber se o knight ta acordado e devolve
// outro valor booleano indicando se pode realizar o ataque rapido ou não.
func CanFastAttack(knightIsAwake bool) bool {
	return !knightIsAwake
}

// CanSpy recebe tres valores booleanos e retorna um único valor booleando
// que indica se ação de espiar é possível
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	return knightIsAwake || archerIsAwake || prisonerIsAwake
}

/*
CanSignalPrisoner recebe dois valores booleanos sobre o estado de dois
personagens e retorna um valor booleano indicando se poderá ser enviado
um sinal de para o prisioneiro.
*/
func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	return prisonerIsAwake && !archerIsAwake
}

/*
CanFreePrisoner recebe 4 valores booleanos referentes aos personagens e retorna um único
valor booleano indicando se poderá resgatar o amigo prisioneiro.
*/
func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	return (prisonerIsAwake && !knightIsAwake && !archerIsAwake) || (petDogIsPresent && !archerIsAwake)
}

func main() {

	var (
		knightIsAwake   = false
		archerIsAwake   = false
		prisonerIsAwake = false
		petDogIsPresent = true
	)

	fmt.Println(CanFastAttack(knightIsAwake))
	fmt.Println(CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake))
	fmt.Println(CanSignalPrisoner(archerIsAwake, prisonerIsAwake))
	fmt.Println(CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent))

}
