package main

import (
	"fmt"
)

func trap1(h []int) int {
	totalWater := 0
	for p := 0; p < len(h); p++ {
		left, right, maxLeft, maxRight := p, p, 0, 0
		for left >= 0 {
			maxLeft = max(maxLeft, h[left])
			left--
		}
		for right < len(h) {
			maxRight = max(maxRight, h[right])
			right++
		}
		currentWater := min(maxLeft, maxRight) - h[p]
		if currentWater >= 0 {
			totalWater += currentWater
		}
	}
	return totalWater
}

func trap2(h []int) int {
	totalWater := 0
	left, right, maxLeft, maxRight := 0, len(h)-1, 0, 0
	for left < right {
		if h[left] < h[right] {
			if h[left] > maxLeft {
				maxLeft = h[left]
			} else {
				totalWater += maxLeft - h[left]
			}
			left++
		} else {
			if h[right] > maxRight {
				maxRight = h[right]
			} else {
				totalWater += maxRight - h[right]
			}
			right--
		}
	}
	return totalWater
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(trap2([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
