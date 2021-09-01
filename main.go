package main

import (
	"fmt"
	"pkg/bank/types"
	"pkg/bank/withdraw"
)

func main() {
	var pointer *int64
	var amount int64 = 50_000_00
	pointer = &amount

	fmt.Println(pointer) // address
	fmt.Println(*pointer) // value
}