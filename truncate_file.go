package main

import (
	"log"
	"os"
)

func main() {
	err := os.Truncate("empty.txt", 100)
	if err != nil {
		log.Fatal(err)
	}
}
