package main

import "sort"

func searchRange(nums []int, target int) []int {
	// 查第一个大于等于8的 --> 3
	a := sort.SearchInts(nums, target)
	if a == len(nums) || nums[a] != target {
		return []int{-1, -1}
	}
	// b := sort.SearchInts(nums,target+1)-1
	b := sort.Search(len(nums), func(i int) bool {
		return nums[i] > 8
	}) - 1
	return []int{a, b}
}
