package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// 在系统的临时文件目录创建一个临时目录(linux系统默认的是/tmp)
	tempDirPath, err := ioutil.TempDir("", "myTempDir")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("TempDir created:", tempDirPath)

	// 在刚创建的临时目录里面创建一个临时文件
	tempFile, err := ioutil.TempFile(tempDirPath, "myTempFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("TempFile created:", tempFile.Name())

	// 关闭临时文件
	err = tempFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	// 删除之前创建的临时文件和临时目录
	err = os.Remove(tempFile.Name())
	if err != nil {
		log.Fatal(err)
	}

	err = os.Remove(tempDirPath)
	if err != nil {
		log.Fatal(err)
	}
}
