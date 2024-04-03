package main

import "fmt"

// Implantacao do algoritmo de arvore binaria
func ArvoreBinaria(lista []int, item int) int {
	baixo := 0
	alto := len(lista) - 1
	var chute int
	var meio int

	for baixo <= alto {
		meio = (alto + baixo) / 2
		chute = lista[meio]

		if chute == item {
			return item
		}

		if chute > item {
			alto = meio - 1
		} else {
			baixo = meio + 1
		}
	}
	return -1
}

func main() {
	fila := []int{1, 2, 7, 10, 15, 17, 21, 23, 30}
	item := 2

	if ArvoreBinaria(fila, item) != -1 {
		fmt.Printf("Number %d found at the list!", item)
	} else {
		fmt.Println("Number missing, nao encontrado")
	}
}
