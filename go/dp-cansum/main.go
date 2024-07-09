package main

import (
	"fmt"
)

// func canSum(target int, nums []int) bool {
// 	var dfs func(int) bool
// 	f := make(map[int]bool, target+1)
// 	dfs = func(sum int) bool {
// 		if _, ok := f[sum]; ok {
// 			return f[sum]
// 		}
// 		if sum == 0 {
// 			return true
// 		}
// 		if sum < 0 {
// 			return false
// 		}
// 		for i := range nums {
// 			if dfs(sum - nums[i]) {
// 				f[sum] = true
// 				return true
// 			}
// 		}

// 		return f[sum]
// 	}

// 	return dfs(target)

// }

// func howSum(target int, nums []int) []int {
// 	var dfs func(int) []int
// 	f := make(map[int][]int, target+1)
// 	for i := range f {
// 		f[i] = make([]int, target)
// 	}
// 	dfs = func(sum int) []int {
// 		if sum == 0 {
// 			return []int{}
// 		}
// 		if sum < 0 {
// 			return nil
// 		}
// 		if _, ok := f[sum]; ok {
// 			return f[sum]
// 		}
// 		for i := range nums {
// 			res := dfs(sum - nums[i])
// 			if res != nil {
// 				if len(append(res, nums[i])) < len(f[sum]) {
// 					f[sum] = append(res, nums[i])
// 					return f[sum]
// 				}

// 			}
// 		}
// 		f[sum] = nil
// 		return nil
// 	}
// 	return dfs(target)
// }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func canSumTable(target int, nums []int) bool {
	f := make(map[int]bool, target+1)
	f[0] = true
	for i := 0; i <= target; i++ {
		for _, num := range nums {
			if f[i] {
				f[i+num] = true
			}
		}
	}
	return f[target]
}

func main() {
	// fmt.Println(howSum(8, []int{2, 3, 5}))
	fmt.Println(canSumTable(11, []int{2, 4, 6}))
}
