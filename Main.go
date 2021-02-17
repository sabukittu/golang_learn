package main

import (
	"fmt"
)

func main() {
	for x := 65; x < 91; x++ {
		fmt.Printf("ASCII Char for No %v = %#U\n", x, x)
	}
}
