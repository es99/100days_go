package main

import (
	"testing"
)

func TestOla(t *testing.T) {

	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado: %s\nesperado: %s\n", resultado, esperado)
		}
	}

	t.Run("diz olá para as pessoas", func(t *testing.T) {
		resultado := Ola("Chris", "")
		esperado := "Hello, Chris"

		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diz 'Olá, mundo' quando uma string vazia for passada", func(t *testing.T) {
		resultado := Ola("", "")
		esperado := "Hello, world!"

		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("em espanhol", func(t *testing.T) {
		resultado := Ola("Engels", "espanhol")
		esperado := "Hola, Engels"
		verificaMensagemCorreta(t, resultado, esperado)
	})

}

func TestMedia(t *testing.T) {
	resultado := Media(100, 100, 100)
	esperado := 100

	if float64(resultado) != float64(esperado) {
		t.Errorf("resultado %v, esperado %v", resultado, esperado)
	}
}
