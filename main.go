package main

import "fmt"

func main() {
	names := []string{"Hyejin", "Jisang", "Hangyeol"}
	names = append(names, "Younggook")
	fmt.Println(names)
}
