package main

import (
	"app/internal"
	"os"
)

func main() {
	err := internal.Host()
	if err != nil {
		os.Exit(1)
	}
}
