// Apenas para testar variáveis e o básico sobre o package net/http
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)

	// Inicia o servidor HTTP na porta 8080
	fmt.Println("Servidor iniciado na porta 8080")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	fmt.Fprintf(w, "Method: %q\nURL.Path: %q\n", r.Method, r.URL.Path)

	circles := queryParams.Get("circles")

	num, err := strconv.Atoi(circles)
	if err != nil {
		fmt.Fprintf(w, "Erro na conversao da string para inteiro %s", err)
	} else {
		fmt.Printf("numero convertido para inteiro %d", num)
	}

	if circles != "" {
		fmt.Fprintf(w, "O valor do parametro circles: %s", circles)
	} else {
		fmt.Fprintf(w, "O parametro circles n foi fornecido na consulta URL")
	}
}
