package main

import (
	"fmt"
)

func main() {
	st := "amqwma"

	for i, j := 0, len(st)-1; i < j; {

		x := string(st[i])
		fmt.Println(x)
		if st[i] != st[j] {
			fmt.Println("false")
			return
		}
		i++
		j--
	}
	fmt.Println("true")
}
