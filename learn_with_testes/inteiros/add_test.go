package inteiros

import (
	"testing"
)

func TestAdd(t *testing.T) {
	soma := Add(2, 2)
	esperado := 4

	if soma != esperado {
		t.Errorf("esperado: %d\nresultado: %d", esperado, soma)
	}
}
