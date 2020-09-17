package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"Pizza", "Chicken", "Kimchi"}
	redo := person{name: "redo", age: 33, favFood: favFood}
	fmt.Println(redo.name)
}
