package main

import (
	"fmt"
)

func main() {
	var empty []int

	alphas := []string{"abc", "def", "ghi", "jkl"}

	empty = append(empty, 123)
	empty = append(empty, 456)
	fmt.Println(empty)

	alphas = append(alphas, "pqr", "stu")
	fmt.Println(alphas)
}
