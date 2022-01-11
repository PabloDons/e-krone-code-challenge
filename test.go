package main

import (
	"fmt"
)

func main() {
	generated := generate(false)
	fmt.Println(generated)
	fmt.Println(testValidity(generated))
	fmt.Println(testValidity("123-abc"))
}
