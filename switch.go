package main

import (
	"fmt"
)

func main() {
	var num int = 2

	switch {
	case num == 1:
		fmt.Println("One")
	case num == 2:
		fmt.Println("Two")
	case num >= 2:
		fmt.Println("Many")
	default:
		fmt.Println("Thrown over boat")
	}
}
