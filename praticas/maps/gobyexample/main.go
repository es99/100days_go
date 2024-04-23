package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v3 := m["k3"]
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	clear(m)
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}

	n3 := map[string]int{"idade": 36}
	if maps.Equal(n2, n3) {
		fmt.Println("n2 == n3")
	} else {
		fmt.Println("n2 != n3")
	}

	nomes := map[string]string{"nome": "Engels", "sexo": "masc"}
	v, prs := nomes["nacionalidade"]
	if prs {
		fmt.Printf("A chave esta presente, com o valor %s", v)
	} else {
		fmt.Println("A chave n esta presente")
	}
}
