package main

import (
	"fmt"
	"math/rand"
)

func Shuffle(sequencia []string) []string {
	rand.Shuffle(len(sequencia), func(i, j int) {
		sequencia[i], sequencia[j] = sequencia[j], sequencia[i]
	})
	return sequencia
}

func main() {
	letras := []string{"a", "b", "c", "d", "e", "f"}
	for i := 0; i < 5; i++ {
		fmt.Println(Shuffle(letras))
	}
}
