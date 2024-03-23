package main

import "fmt"

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100
}

func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
}

func CalculateCost(carsCount int) uint {
	return uint((carsCount/10)*95000 + (carsCount%10)*10000)
}

func main() {

	var (
		carrosProduzidos int     = 1105
		taxaSucesso      float64 = 90
		numeroDeCarros   int     = 21
	)

	fmt.Println("carros produzidos por hora:", CalculateWorkingCarsPerHour(carrosProduzidos, taxaSucesso))
	fmt.Println("carros produzidos por minuto:", CalculateWorkingCarsPerMinute(carrosProduzidos, taxaSucesso))
	fmt.Println("custo dos carros:", CalculateCost(numeroDeCarros))
}
