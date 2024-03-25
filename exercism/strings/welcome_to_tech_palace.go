package main

import (
	"fmt"
	"strings"
)

// WelcomeMessage recebe um string com o nome do cliente e retorna com uma mensagem
// de boas-vindas contendo o nome do cliente.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	return strings.Repeat("*", numStarsPerLine) + "\n" + welcomeMsg + "\n" + strings.Repeat("*", numStarsPerLine)
}

func CleanupMessage(oldMsg string) string {
	return strings.TrimSpace(strings.ReplaceAll(oldMsg, "*", ""))
}

func main() {
	message := `
	**************************
	*    BUY NOW, SAVE 10%   *
	**************************
	`

	customer := "Engels"
	numStars := 10
	welcomeMessage := "Sejam bem-vindo"

	fmt.Println(WelcomeMessage(customer))
	fmt.Println(AddBorder(welcomeMessage, numStars))
	fmt.Println(CleanupMessage(message))
}
