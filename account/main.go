package main

import (
	"fmt"

	"github.com/duqrldudgns/Go/account/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	fmt.Println(account)
}
