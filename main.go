package main

import (
	"fmt"
)

func main() {
	num := 4
	for num < 15 {
		if num > 0 {
			fmt.Println(num)
			num++
		}
	}
}
