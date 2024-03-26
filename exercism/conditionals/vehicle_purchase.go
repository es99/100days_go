package main

import "fmt"

// NeedsLicence determines whether a license is needed to drive a type of vehicle.
// 'car' and 'truck' needs a licence, otherwise diferents vehicle don't need a license.
func NeedsLicence(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	} else {
		return false
	}
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle
// that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	msg := " is clearly the better choice."

	if option1 < option2 {
		return option1 + msg
	} else {
		return option2 + msg
	}
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		return originalPrice * 0.8
	} else if age >= 10 {
		return originalPrice * 0.5
	} else {
		return originalPrice * 0.7
	}
}

func main() {
	tipo := "bike"
	idade := 2

	fmt.Println(NeedsLicence(tipo))
	fmt.Println(ChooseVehicle("Corsa", "Clio"))
	fmt.Println(CalculateResellPrice(1000, float64(idade)))
}
