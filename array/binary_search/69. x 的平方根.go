package main

// 相当于在[0,x]中
func mySqrt(x int) int {
	l, r := 0, x
	for l <= r {
		mid := (l + r) / 2
		if mid*mid > x {
			r = mid - 1
		} else if mid*mid < x {
			l = mid + 1
		} else {
			return mid
		}
	}
	return r
}
