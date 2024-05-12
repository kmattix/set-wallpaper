package main

import (
	"log"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		log.Fatal("No arguments provided.")
	} else if len(args) > 2 {
		log.Fatal("Too many arguments provided.")
	}
}
