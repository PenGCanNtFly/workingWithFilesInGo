package main

import (
	"io"
	"log"
	"os"
)

func main() {
	srcFile, err := os.Open("copySrcFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer srcFile.Close()

	dstFile, err := os.Create("copyDstFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer dstFile.Close()

	bytesWritten, err := io.Copy(dstFile, srcFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes.", bytesWritten)

	err = dstFile.Sync()
	if err != nil {
		log.Fatal(err)
	}
}
