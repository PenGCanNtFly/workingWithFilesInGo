/*
* `mv empty.txt empty.log`
 */
package main

import (
	"log"
	"os"
)

func main() {
	oldpath := "empty.txt"
	newpath := "empty.log"
	err := os.Rename(oldpath, newpath)
	if err != nil {
		log.Fatal(err)
	}
}
