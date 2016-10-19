package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Open("empty.txt")
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	file, err = os.OpenFile("empty.txt", os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
}
