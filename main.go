package main

import (
	"fmt"

	"github.com/redonearth/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("Redo")
	account.Deposit(10)
	fmt.Println(account.Balance())
}
