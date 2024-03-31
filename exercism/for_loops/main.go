package main

import "fmt"

func TotalBirdCount(birdsPerDay []int) int {
	var birds int
	for i := 0; i < len(birdsPerDay); i++ {
		birds += birdsPerDay[i]
	}
	return birds
}

func BirdsInWeek(birdsPerDay []int, week int) int {
	/* meu codigo antigo
	initial := (week * 7) - 7
	semana := birdsPerDay[initial:]
	sum := 0
	for i := 0; i < 7; i++ {
		sum += semana[i]
	}
	return sum
	*/
	var birds int
	for i := (week * 7) - 7; i < week*7; i++ {
		birds += birdsPerDay[i]
	}
	return birds
}

func FixBirdCountLog(birdsPerDay []int) []int {
	/* meu codigo antigo
	for i := range birdsPerDay {
		if i%2 == 0 {
			birdsPerDay[i] += 1
		}
		continue
	}
	return birdsPerDay
	*/
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i]++
	}
	return birdsPerDay
}

func main() {
	birdsPerDay := []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}
	passaros := []int{2, 5, 0, 7, 4, 1}

	fmt.Println(TotalBirdCount(birdsPerDay))
	fmt.Println(BirdsInWeek(birdsPerDay, 2))
	fmt.Println(FixBirdCountLog(passaros))
}
