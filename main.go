package main

import (
	"fmt"
)

func main() {
	a := 2
	b := &a
	*b = 2020
	fmt.Println(a)
}
