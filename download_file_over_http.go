package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	downloadFile, err := os.Create("baidu.htm")
	if err != nil {
		log.Fatal(err)
	}
	defer downloadFile.Close()

	url := "http://www.baidu.com"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	numBytesWritten, err := io.Copy(downloadFile, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Downloaded %d bytes file", numBytesWritten)
}
