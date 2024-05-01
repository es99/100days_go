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
