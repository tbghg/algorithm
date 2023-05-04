package double_pointer

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	idx := len(nums) - 1
	l, r := 0, len(nums)-1
	for i, v := range nums {
		nums[i] = v * v
	}
	for idx >= 0 {
		if nums[l] < nums[r] {
			res[idx] = nums[r]
			r--
		} else {
			res[idx] = nums[l]
			l++
		}
		idx--
	}
	return res
}

/*
func sortedSquares(nums []int) []int {
    res :=make([]int,len(nums))
    idx := len(nums)-1
    l,r := 0,len(nums)-1
    for idx >= 0 {
        a,b := nums[l]*nums[l],nums[r]*nums[r]
        if a < b {
            res[idx] = b
            r--
        } else {
            res[idx] = a
            l++
        }
        idx--
    }
    return res
}
*/
