package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "Kishor"
	//fmt.Print(str)
	strlen := ""
	for i := 0; i < len(str); i++ {
		if !strings.Contains(strlen, string(str[i])) {
			count := 1
			for j := i + 1; j < len(str); j++ {
				if str[i] == str[j] {
					count++
				}
			}
			strlen = strlen + " " + string(str[i]) + ":" + strconv.Itoa(count)
		}
	}
	fmt.Println(strlen)

}
