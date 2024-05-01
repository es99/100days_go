package main

import (
	"testing"
)

func TestOla(t *testing.T) {
	resultado := Ola("Chris")
	esperado := "Hello, Chris"

	if resultado != esperado {
		t.Errorf("resultado: %s, esperado: %s", resultado, esperado)
	}
}

func TestMedia(t *testing.T) {
	resultado := Media(100, 100, 100)
	esperado := 100

	if float64(resultado) != float64(esperado) {
		t.Errorf("resultado %v, esperado %v", resultado, esperado)
	}
}
