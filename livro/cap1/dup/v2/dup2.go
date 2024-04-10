// Dup2 exibe a contagem e o texto das linhas que aparecem mais de uma
// vez na entrada. Ele le stdin ou uma lista de arquivos nomeados.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(f *os.File, counts map[string]int, filenames map[string]bool, nome string) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			filenames[nome] = true
		}
	}
}

func countLinesStdin(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()]++
	}
}

func main() {
	counts := make(map[string]int)
	filenames := make(map[string]bool)
	files := os.Args[1:]

	if len(files) == 0 {
		countLinesStdin(os.Stdin, counts)
	} else {
		for _, v := range files {
			f, err := os.Open(v)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, filenames, v)
			f.Close()
		}
	}

	if len(files) == 0 {
		fmt.Println("Entrada padrao: teclado")

		// codigo duplicado, resolver no futuro
		for k, v := range counts {
			if v > 1 {
				fmt.Printf("%d\t%s\t\n", v, k)
			}
		}
	} else {
		fmt.Println("Arquivos onde ocorrem duplicadas:")
		for k := range filenames {
			fmt.Printf("%s ", k)
			fmt.Printf("\n")
		}

		// codigo duplicado, resolver no futuro
		for k, v := range counts {
			if v > 1 {
				fmt.Printf("%d\t%s\t\n", v, k)
			}
		}
	}

}
