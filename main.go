package main

import (
	"fmt"

	"github.com/lechnerc77/gotestdemo/calculator"
)

func main() {
	sum := calculator.Add(23, 42)
	fmt.Println(sum)
}
