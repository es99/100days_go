// Programa que recebe le arquivos nomeados como argumentos da linha de comando
// e procura nestes arquivos a palavra 'amor', ao final do programa, ele exibe na tela
// o nome do arquivo, a quantidade de palavras 'amor' encontradas e o texto que estava
// no arquivo.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countLines(f *os.File, nome string) {
	input := bufio.NewScanner(f)
	linhas := []string{}
	countAmor := 0

	for input.Scan() {
		linhas = append(linhas, input.Text())
		linha := strings.Split(input.Text(), " ")
		for i := 0; i < len(linha); i++ {
			if linha[i] == "amor" {
				countAmor++
			}
		}
	}
	fmt.Printf("Num de %d palavras encontradas no arquivo %s:\n ", countAmor, nome)
	for i := 0; i < len(linhas); i++ {
		fmt.Println(linhas[i])
	}
}

func main() {
	files := os.Args[1:]

	if len(files) == 0 {
		fmt.Println("falta passar argumentos. programa encerrado.")
	} else {
		for _, v := range files {
			f, err := os.Open(v)
			if err != nil {
				fmt.Fprintf(os.Stderr, "procura_amor: %v\n", err)
			}
			countLines(f, v)
			f.Close()
		}
	}
}
