package main

import (
	"fmt"
	"sort"
)

func binarySearch(a []int, target int) int {
	l, r := 0, len(a)-1
	for l <= r {
		mid := (l + r) / 2
		if a[mid] == target {
			return mid
		}
		if a[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}

func main() {
	// 二分查找要求数组有序
	a := []int{1, 4, 56, 324, 978, 3465, 5355}
	// 返回 [0， n) 中 f(i) 为真的最小索引 i
	fmt.Print(sort.Search(len(a), func(i int) bool {
		return a[i] >= 10
	}))
	//fmt.Print(binarySearch(a, 10))
}
