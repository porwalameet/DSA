package main

import (
	"fmt"
)

func canConstruct(target string, wordBank []string) bool {
	f := make(map[string]bool)
	wordSet := make(map[string]bool)
	for _, word := range wordBank {
		wordSet[word] = true
	}
	return dfs(target, wordSet, f)

}
func dfs(s string, wordSet, f map[string]bool) bool {
	if value, exists := f[s]; exists {
		return value
	}
	if _, exists := wordSet[s]; exists {
		return true
	}
	for i := 1; i < len(s); i++ {
		prefix := s[:i]
		if wordSet[prefix] && dfs(s[i:], wordSet, f) {
			f[s] = true
			return true
		}
	}
	f[s] = false
	return false
}

func main() {
	fmt.Println(canConstruct("abcdef", []string{"ab", "abc", "ef", "def", "abcd"}))
	fmt.Println(canConstruct("eeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeeee", "abcd"}))
}
