package main

import (
	"api/pkg"
	"os"
)

func main() {
	err := pkg.Host()
	if err != nil {
		os.Exit(1)
	}
}
