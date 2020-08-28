package main

import "fmt"

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	repeatMe("redo", "reco", "jisang", "hyejin", "nico")
}
