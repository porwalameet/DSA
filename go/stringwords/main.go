package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "This, that, and the other."
	replacer := strings.NewReplacer(",", "", ";", "", ".", "")
	s = replacer.Replace(s)
	// s = strings.ReplaceAll(s, ",", "")
	words := strings.Fields(s)
	for i, w := range words {
		fmt.Printf("%d-->%s\n", i, w)
	}
	fmt.Println()
}
