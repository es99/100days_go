package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	s, sep := "", ""

	for _, v := range os.Args[1:] {
		s += sep + v
		sep = " "
	}
	fmt.Println(s)
	fmt.Printf("%vs elapsed\n", time.Since(start).Microseconds())
}
