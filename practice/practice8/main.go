package main

import "fmt"

func main() {
	target := twoNumberSumTarget([]int{1, 2, 4, 64, 64, 323}, 128)
	fmt.Println(target)
}

func twoNumberSumTarget(nums []int, target int) []int {
	m := make(map[int]int, 10)
	for i := 0; i < len(nums); i++ {
		if p, ok := m[target-nums[i]]; ok {
			return []int{i, p}
		}
		m[nums[i]] = i
	}
	return []int{}
}
