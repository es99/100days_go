// Le de um arquivo linhas correspondentes a emails e gera uma lista sem linhas em branco
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("lista_emails.txt")
	if err != nil {
		fmt.Printf("Ocorreu um erro ao abrir o arquivo %v\n", f)
	}
	erro := ExcluiLinhasEmBranco(f)
	if erro != nil {
		fmt.Println("Houve algum erro ao gerar o arquivo:", erro)
	}
}

func ExcluiLinhasEmBranco(f *os.File) error {
	input := bufio.NewScanner(f)
	arquivo, err := os.Create("lista_emails-tratada.txt")
	if err != nil {
		return err
	}
	defer arquivo.Close()

	for input.Scan() {
		if input.Text() == "" || input.Text() == " " {
			continue
		}
		arquivo.WriteString(input.Text() + "\n")
	}
	fmt.Println("Dados foram escritos com sucesso!")
	return nil
}
