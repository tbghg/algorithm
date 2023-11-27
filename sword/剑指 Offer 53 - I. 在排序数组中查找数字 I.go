package sword

import "sort"

func search(nums []int, target int) int {
	f := binarySearch(nums, target)
	e := binarySearch(nums, target+1)
	return e - f
}

func binarySearch(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}

// 使用sort.Search
func search1(nums []int, target int) int {

	f := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})

	e := sort.Search(len(nums), func(i int) bool {
		return nums[i] > target
	})

	return e - f
}
