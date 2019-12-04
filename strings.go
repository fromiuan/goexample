package main

import (
	"log"
	"strings"
)

func main() {
	sentence := "I'm a sentence made up of words"
	fields := strings.Fields(sentence)
	log.Println(fields)
}
