// Fetch exibe o conteudo encontrado em cada URL especificada.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	prefixo := "http://"
	for _, url := range os.Args[1:] {

		if !strings.HasPrefix(url, prefixo) {
			url = prefixo + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := io.Copy(os.Stdout, resp.Body)
		status := resp.Status
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("\nStatus code - %s\n\n", status)
		fmt.Printf("%v\n", b)
	}
}
