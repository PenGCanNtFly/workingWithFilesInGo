package main

import (
	"compress/gzip"
	"io"
	"log"
	"os"
)

func main() {
	gzipFile, err := os.Open("test.txt.gz")
	if err != nil {
		log.Fatal(err)
	}

	gzipReader, err := gzip.NewReader(gzipFile)
	if err != nil {
		log.Fatal(err)
	}

	outputFile, err := os.Create("unzipped.txt")
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(outputFile, gzipReader)
	if err != nil {
		log.Fatal(err)
	}
}
