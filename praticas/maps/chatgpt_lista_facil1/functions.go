/*
5 functions de nivel easy para praticar o conceito de maps
*/

package chatgptlistafacil1

import (
	"fmt"
)

// AdicionaElemento recebe um map como entrada, uma chave e um valor e adiciona o par chave-valor ao map.
func AdicionaElemento(dados map[string]float64, key string, value float64) {
	_, status := dados[key]
	if status {
		fmt.Printf("%s ja existe no dicionario\n", key)
	} else {
		dados[key] = value
		fmt.Printf("%s adicionado no dicionario com o valor %f\n", key, value)
	}
}

// ContaCaractere recebe uma string como entrada e retorna um map contendo a contagem de cada caractere na string.
func ContaCaractere(palavra string) map[rune]int {
	contagem := make(map[rune]int)
	for _, c := range palavra {
		contagem[c] += 1
	}
	return contagem
}

// VerificaValor recebe um map como entrada e um valor como parâmetro,
// e retorne verdadeiro se o valor estiver presente no map e falso caso contrário.
func VerificaValor(mapa map[string]string, valor string) bool {
	for _, v := range mapa {
		if v == valor {
			return true
		}
	}
	return false
}

// VerificaChave recebe um map como entrada e uma chave, e retorna verdadeiro se a chave estiver presente
// no map e falso caso contrário.
func VerificaChave(contatos map[string]int, chave string) bool {
	_, status := contatos[chave]
	if status {
		return true
	} else {
		return false
	}
}

// Escreva uma função que receba um map como entrada e uma chave, e remova o par chave-valor correspondente do map, se existir.
func RemoveElemento(mapa map[string]int, chave string) {
	_, status := mapa[chave]
	if status {
		delete(mapa, chave)
		fmt.Printf("%s deletada", chave)
	} else {
		fmt.Println("Esta chave n existe no mapa.")
	}
}
