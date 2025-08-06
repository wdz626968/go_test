package main

import "fmt"

// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
// 找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
// 例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
func main() {
	nums := []int{4, 1, 2, 1, 2}
	nums1 := []int{1}
	fmt.Println(isSingleNumber(nums))
	fmt.Println(isSingleNumber(nums1))
}

func isSingleNumber(nums []int) int {
	var digitMap map[int]int = make(map[int]int)

	for _, v := range nums {
		digitMap[v] += 1
	}

	for k, v := range digitMap {
		if v == 1 {
			return k
		}
	}
	return 0
}
