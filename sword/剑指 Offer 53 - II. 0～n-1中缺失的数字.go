package sword

import "sort"

// 使用标准库
func missingNumber1(nums []int) int {
	return sort.Search(len(nums), func(i int) bool {
		return nums[i] != i
	})
}

func missingNumber(nums []int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == mid {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
