package main

import (
	"fmt"

	"github.com/redonearth/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("Redonearth")
	fmt.Println(account)
}
