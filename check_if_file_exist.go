package main

import (
	"log"
	"os"
)

func main() {
	fileInfo, err := os.Stat("empty.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist.")
		}
	}
	log.Println("File does exist. File information:")
	log.Println(fileInfo)
}
