/*
* `rm empty.txt`
 */
package main

import (
	"log"
	"os"
)

func main() {
	err := os.Remove("empty.log")
	if err != nil {
		log.Fatal(err)
	}

}
