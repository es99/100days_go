package main

import (
	"testing"
)

func TestOla(t *testing.T) {
	resultado := Ola()
	esperado := "Hello, world"

	if resultado != esperado {
		t.Errorf("resultado: %s, esperado: %s", resultado, esperado)
	}
}
