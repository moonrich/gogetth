package main

import (
	"fmt"

	"github.com/merxer/demo/bank"
)

func init() {
	fmt.Println("Initial")
}

func main() {
	println("Hello, " + bank.Title)
}
