package main

import (
	"os"

	"github.com/evgenivanovi/gophkeeper/internal/client/boot"
)

func main() {
	err := boot.RootCMD.Execute()
	if err != nil {
		os.Exit(1)
	}
}
