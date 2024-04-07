// Dup2 exibe a contagem e o texto das linhas que aparecem mais de uma
// vez na entrada. Ele le stdin ou uma lista de arquivos nomeados.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()]++
	}
}

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, v := range files {
			f, err := os.Open(v)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for k, v := range counts {
		if v > 1 {
			fmt.Printf("%d\t%s\n", v, k)
		}
	}

}
