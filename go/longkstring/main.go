package main

import "fmt"

func longestKString(s string, k int) int {
	counter := map[byte]int{}
	j := 0
	for i := range s {
		counter[s[i]]++
		if len(counter) > k {
			counter[s[j]]--
			if counter[s[j]] == 0 {
				delete(counter, s[j])
			}
			j++
		}
	}
	return len(s) - j
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "aabbcc"
	k := 3
	fmt.Println(longestKString(s, k))
}
