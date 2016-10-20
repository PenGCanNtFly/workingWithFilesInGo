package main

import (
	"io/ioutil"
	"log"
)

func main() {
	err := ioutil.WriteFile("quickWriteToFile.txt", []byte("Quick Write to File\n"), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
