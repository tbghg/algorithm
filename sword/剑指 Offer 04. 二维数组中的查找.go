package sword

import "sort"

// 方法一：暴力
// 方法二：遍历行，然后每一行二分查找一下，看看有没有（算半个暴力）
func findNumberIn2DArray(matrix [][]int, target int) bool {
	for _, array := range matrix {
		i := sort.SearchInts(array, target)
		if i < len(array) && array[i] == target {
			return true
		}
	}
	return false
}
