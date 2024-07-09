package main

import (
	"fmt"
	"strings"
)

// func allConstruct(target string, wordBank []string) [][]string {
// 	array := make([][]string, 0)

// 	for i := range array {
// 		array[i] = make([]string, 0)
// 	}

// 	var dfs(s string) [][]string {
// 	if len(s) == 0 {
// 		return [][]string{}
// 	}
// 	targetWays := make([]string, 0)
// 	for _, word := range wordBank {
// 		if strings.Index(s, word) == 0 {
// 			ways := allConstruct(strings.SplitAfter(s, word)[0])
// 			//array[0] = append([]string{word}, array[]...)
// 			targetWays = append(ways..., targetWays)
// 		}
// 	}
// 	return array
// 	}
// 	return dfs(target)
// }

func countWays(s string, elements []string) int {
	//if s == "" {
	if len(s) == 0 {
		return 1
	}
	totalCount := 0
	for _, ele := range elements {
		if strings.Index(s, ele) == 0 {
			numWays := countWays(s[len(ele):], elements)
			totalCount += numWays
		}
	}
	return totalCount
}

func main() {
	//fmt.Println(countWays("abcdef", []string{"abc", "ab", "cd", "ef", "def", "c"}))
	fmt.Println(countWays("practice", []string{"p", "h", "he", "ra", "ti", "c", "ce", "ac", "pr", "prac", "act", "ice", "tice", "ice"}))
}
