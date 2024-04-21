package main

import (
	"bufio"
	"conversor"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	valores := []int{}
	if len(os.Args[1:]) > 0 {
		for _, value := range os.Args[1:] {
			valor, err := strconv.ParseInt(value, 10, 0)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Erro: %g", err)
			}
			d := conversor.Dia(valor)
			h := conversor.Hora(valor)
			Imprime(d, h)
		}
	} else {
		for input.Scan() {
			inteiro, err := strconv.ParseInt(input.Text(), 10, 0)
			if err != nil {
				fmt.Fprintf(os.Stderr, "impossivel converter o numero: %g", err)
			}
			valores = append(valores, int(inteiro))
		}
		for _, v := range valores {
			d := conversor.Dia(v)
			h := conversor.Hora(v)
			Imprime(d, h)
		}
	}

}

func Imprime(d conversor.Dia, h conversor.Hora) {
	fmt.Printf("%s = %s, %s = %s\n", d, conversor.DToH(d), h, conversor.HToD(h))
}
