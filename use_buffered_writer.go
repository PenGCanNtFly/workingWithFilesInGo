package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("bufio.txt", os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bufferedWriter := bufio.NewWriter(file)

	// Write bytes to buffer
	bytesWritten, err := bufferedWriter.Write([]byte{65, 66, 67})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes Written: %d\n", bytesWritten)

	// Write string to buffer
	bytesWritten, err = bufferedWriter.WriteString("Buffered string\n")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Bytes Written: %d\n", bytesWritten)

	// Check how much is stored in buffer waiting
	unflushedBufferSize := bufferedWriter.Buffered()
	log.Printf("Bytes buffered: %d\n", unflushedBufferSize)

	// Check how much buffer is avaliable
	bytesAvaliable := bufferedWriter.Available()
	log.Printf("Bytes avaliable: %d\n", bytesAvaliable)

	// Write memory buffer to disk
	err = bufferedWriter.Flush()
	if err != nil {
		log.Fatal(err)
	}

	// ???
	bufferedWriter.Reset(bufferedWriter)

	// See how much buffer is avaliable
	bytesAvaliable = bufferedWriter.Available()
	log.Printf("Bytes avaliable: %d\n", bytesAvaliable)

	// Resize buffer
	bufferedWriter = bufio.NewWriterSize(bufferedWriter, 8000)

	// Check how much buffer is avaliable after resize
	bytesAvaliable = bufferedWriter.Available()
	log.Printf("Bytes avaliable: %d\n", bytesAvaliable)
}
