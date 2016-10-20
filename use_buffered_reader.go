package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.Open("use_buffered_reader.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bufferedReader := bufio.NewReader(file) // bufio.NewReader返回*Reader类型，Reader类型有Peek|Read|ReadByte|ReadByte|ReadBytes|ReadString等方法
	byteSlice := make([]byte, 5)

	// 不移动reader指针，返回后 n bytes
	byteSlice, err = bufferedReader.Peek(5)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Peeked at 5 bytes: %d\n", byteSlice)

	//　读取并移动指针
	numBytesRead, err := bufferedReader.Read(byteSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Read %d bytes into byteSilce: %s\n", numBytesRead, byteSlice)

	// read and return a single byte
	mByte, err := bufferedReader.ReadByte()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Read a single byte: %c\n", mByte)

	// 返回直到第一个分隔符之间内容(Read up to and including delimiter, Returns byte slice)
	dataBytes, err := bufferedReader.ReadBytes('\n')
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Read bytes: %s\n", dataBytes)

	// Read up to and including delimiter ,Returns string
	dataString, err := bufferedReader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Read string: %s\n", dataString)

}
