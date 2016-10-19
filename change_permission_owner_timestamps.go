/*
*"chmod 777 empty.txt"
*"chown `id -u`:`id -g` empty.txt"
*""
 */
package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	err := os.Chmod("empty.txt", 0777)
	if err != nil {
		log.Println(err)
	}

	err = os.Chown("empty.txt", os.Getuid(), os.Getgid())
	if err != nil {
		log.Println(err)
	}

	twoDaysFromNow := time.Now().Add(48 * time.Hour)
	fmt.Println(twoDaysFromNow)
	lastAccessTime := twoDaysFromNow
	lastModifiedTime := twoDaysFromNow
	err = os.Chtimes("empty.txt", lastAccessTime, lastModifiedTime)
	if err != nil {
		log.Println(err)
	}
}
