package main

import (
	"fmt"
)

func main() {
	i := 1
	fmt.Println(i - 1)
	for i <= 10 {
		fmt.Println("Let's add 1 more!", i)
		i++
	}
}
