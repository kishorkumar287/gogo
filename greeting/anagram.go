package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	string1 := "Kishor"
	string2 := "Kishoe"
	string1Slice := strings.Split(string1, "")
	string2Slice := strings.Split(string2, "")

	sort.Slice(string1Slice, func(i, j int) bool {
		return string1Slice[i] < string1Slice[j]
	})

	sort.Slice(string2Slice, func(i, j int) bool {
		return string2Slice[i] < string2Slice[j]
	})

	var sortedstring1 string = strings.Join(string1Slice, "")
	var sortedstring2 string = strings.Join(string2Slice, "")

	if sortedstring1 == sortedstring2 {
		fmt.Println("These strings are anagrams")
	} else {
		fmt.Println("These strings not are anagrams")
	}
}
