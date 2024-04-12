// Server4
package main

import (
	"fmt"
	"lissajous"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		queryParams := r.URL.Query()
		cycles := queryParams.Get("cycles")
		if cycles == "" {
			fmt.Fprintf(w, "Nao encontramos nenhum parameto cycles na URL")
		}
		num, err := strconv.Atoi(cycles)
		if err != nil {
			fmt.Fprintf(w, "Erro ao converter formato %q", err)
		} else {
			lissajous.Lissajous(w, num)
		}
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler ecoa o componente Path do URL requisitado
/*
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
*/
