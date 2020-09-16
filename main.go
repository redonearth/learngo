package main

import "fmt"

func main() {
	jisang := map[string]string{"name": "jisang", "age": "33"}
	for _, value := range jisang {
		fmt.Println(value)
	}
}
