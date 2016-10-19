package main

import (
	"log"
	"os"
)

func main() {
	/*
		file, err := os.OpenFile("empty.txt", os.O_WRONLY, 0666)
		if err != nil {
			if os.IsPermission(err) {
				log.Fatal("Error: Write Permission Denied")
			}
		}
		file.Close()
	*/

	file, err := os.OpenFile("empty.txt", os.O_RDONLY, 0666)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("Error: File does not exist")
		} else if os.IsPermission(err) {
			log.Fatal("Error: Read Permission Denied")
		}
	}
	file.Close()
}
