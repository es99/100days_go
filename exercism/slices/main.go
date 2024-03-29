package main

import (
	"fmt"
	"meuprojeto/cards"
)

func main() {
	lista := []int{2, 15, 20, 30, 14}
	fmt.Println("cartas favoritas:", cards.FavoriteCards())
	fmt.Println(cards.GetItem(lista, 3))
	fmt.Println(cards.SetItem(lista, 0, 100))
	fmt.Println(cards.PrependItems(lista))
	fmt.Println(cards.RemoveItem(lista, 3))
}
