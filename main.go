package main

import (
	"fmt"

	"github.com/stevepro7/gotestinglib/adriana"
)

func main() {
	val1 := adriana.Add(2, 3)
	fmt.Println(val1)

	val2 := adriana.Sub(7, 3)
	fmt.Println(val2)
}
