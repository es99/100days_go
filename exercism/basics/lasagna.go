package main

import "fmt"

const OvenTime = 40

func RemainingOvenTime(actualMinutesInOven int) int {
	return OvenTime - actualMinutesInOven
}

func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
}

// ElapsedTime calculates the total time needed to create and bake a lasagna.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return PreparationTime(numberOfLayers) + actualMinutesInOven
}

func main() {
	var tempoAtualDecorrido uint8
	var camadas uint8

	tempoAtualDecorrido = 25
	camadas = 4

	fmt.Println("Tempo restante de forno:", RemainingOvenTime(int(tempoAtualDecorrido)))
	fmt.Println("Tempo de prepacao:", PreparationTime(int(camadas)))
	fmt.Println("Tempo TOTAL decorrido:", ElapsedTime(int(camadas), int(tempoAtualDecorrido)))
}
