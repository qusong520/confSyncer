package main

import (
	"log"
)

func main() {
	log.SetFlags(log.Ldate | log.Lshortfile)
	Execute()
}
