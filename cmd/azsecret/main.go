package main

import (
	"os"
)

func main() {
	if execute() != nil {
		os.Exit(1)
	}
}
