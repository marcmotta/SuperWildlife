// cmd/superwildlife/main.go
package main

import (
	"flag"
	"log"
	"os"

	"superwildlife/internal/superwildlife"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := superwildlife.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
