package main

import (
	"fmt"
	"github.com/progayk/go-programmings-level_12/01/dog"
)

func main() {
	hy := 45
	dy := dog.ConvToDogYears(hy)
	fmt.Printf("%v human years is equal to %v dog years", hy, dy)
}
