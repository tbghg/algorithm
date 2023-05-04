package sliding_window

import "math"

func minSubArrayLen(target int, nums []int) int {
	l, r := 0, 0
	sum, ans := 0, math.MaxInt
	for r < len(nums) {
		sum += nums[r]
		for sum >= target {
			ans = min(ans, r-l+1)
			sum -= nums[l]
			l++
		}
		r++
	}
	if ans == math.MaxInt {
		return 0
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	left,end:=0,0
	sum,res := 0,math.MaxInt32
	for end < n {
		sum += nums[end]
		for sum >= target {
			res = min(res,end-left+1)
			sum-=nums[left]
			left++
		}
		end++
	}
	if res == math.MaxInt32 {
		return 0
	}
	return res
}

func min(a,b int)int{
	if a<b {
		return a
	}
	return b
}
*/
