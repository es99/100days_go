package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()

	for i, v := range os.Args {
		fmt.Printf("Index %d, valor: %s\n", i, v)
	}

	fmt.Printf("%vs elapsed\n", time.Since(start).Microseconds())
}
