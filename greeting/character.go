package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "asdasdasdaafgjopa"
	fmt.Println(a)
	a = strings.Replace(a, "a", "", -1)
	fmt.Println(a)
}
