package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan bool)
	people := [2]string{"Redonearth", "Koongdori"}
	for _, person := range people {
		go isCool(person, channel)
	}
	fmt.Println(<-channel)
	fmt.Println(<-channel)
}

func isCool(person string, channel chan bool) {
	time.Sleep(time.Second * 3)
	fmt.Println(person)
	channel <- true
}
