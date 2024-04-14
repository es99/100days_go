package main

import (
	"fmt"
)

func main() {
	quantities := []float64{1.2, 3.6, 10.5}
	scaledQuantities := ScaleRecipe(quantities, 1)

	fmt.Println(scaledQuantities)
}

func PreparationTime(camadas []string, preparationTimePerLayer int) int {
	if preparationTimePerLayer > 0 {
		return preparationTimePerLayer * len(camadas)
	}
	return len(camadas) * 2
}

func Quantities(camadas []string) (int, float64) {
	var countNoodles int
	var countSauce float64
	for _, value := range camadas {
		if value == "noodles" {
			countNoodles++
		}
		if value == "sauce" {
			countSauce++
		}
	}
	return countNoodles * 50, countSauce * 0.2
}

func AddSecretIngredient(friendList, myList []string) {
	secretIngredient := friendList[len(friendList)-1]
	for i, value := range myList {
		if value == "?" {
			myList[i] = secretIngredient
		}
	}
}

func ScaleRecipe(amounts []float64, portions int) []float64 {
	newAmounts := []float64{}
	for _, value := range amounts {
		if portions == 2 {
			newAmounts = append(newAmounts, value)
		} else if portions < 2 {
			newAmounts = append(newAmounts, value/2.0)
		} else {
			qtd := (value * float64(portions)) / 2
			newAmounts = append(newAmounts, qtd)
		}
	}
	return newAmounts
}
