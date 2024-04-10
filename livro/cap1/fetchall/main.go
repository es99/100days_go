// Fetchall busca URLs em paralelo e informa os tempos gastos e os tamanhos em bytes.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	lista := []string{
		"https://google.com.br",
		"https://youtube.com",
		"https://g.com",
		"https://uol.com.br",
		"https://live.com",
		"ht://mercadolivre.com.br",
		"https://blastingnews.com",
		"https://wikipedia.com",
		"https://ntfl.com",
		"https://instagram.com",
	}

	arquivo, err := os.Create("saida.txt")
	if err != nil {
		fmt.Println("Erro ao criar arquivo:", err)
		return
	}
	defer arquivo.Close()
	os.Stdout = arquivo

	for _, url := range lista {
		go fetch(url, ch)
	}

	for range lista {
		fmt.Println(<-ch) // recebe do canal ch
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // envia para o canal ch
		return
	}

	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("While reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
