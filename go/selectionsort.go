// package main

// import "fmt"

// func main() {
// 	s := []int{5, 4, 7, 3, 6, 19}
// 	for i := 0; i < len(s); i++ {
// 		minPos := i
// 		for j := i + 1; j < len(s); j++ {
// 			if s[j] < s[minPos] {
// 				minPos = j
// 			}
// 		}
// 		t := s[i]
// 		s[i] = s[minPos]
// 		s[minPos] = t
// 	}
// 	fmt.Println(s)
// }
