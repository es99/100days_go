// Dup1 exibe o texto de toda linha que aparece mais de
// uma vez na entrada-padrao, precedida por sua contagem.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	for k, v := range counts {
		if v > 1 {
			fmt.Printf("%d\t%s\n", v, k)
		}
	}
}
