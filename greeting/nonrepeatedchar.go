package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "aaaaad"

	for i := 0; i < len(str); i++ {
		flag := true
		for j := i + 1; j < len(str); j++ {
			if str[i] == str[j] {
				flag = false
				str = strings.Replace(str, string(str[i]), "", -1)
				i--
				break
			}
		}
		//fmt.Print(str)
		if flag {
			fmt.Print(string(str[i]))
			break
		}
	}
}
