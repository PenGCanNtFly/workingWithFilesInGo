package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	err := os.Link("empty.txt", "empty_also.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Creating Symlinks")
	err = os.Symlink("empty.txt", "empty_sym.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileInfo, err := os.Lstat("empty_sym.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Link info: %+v", fileInfo)

	err = os.Lchown("empty_sym.txt", os.Getuid(), os.Getgid())
	if err != nil {
		log.Fatal(err)
	}

}
