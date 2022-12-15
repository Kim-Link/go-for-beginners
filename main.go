package main

import (
	"fmt"

	"github.com/KimLink/go-for-beginners/accounts"
)

func main() {
	accounts := accounts.NewAccount("Link")
	fmt.Println(accounts)
}
