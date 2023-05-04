package main

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			return mid
		}
	}
	return l
}

/*func searchInsert(nums []int, target int) int {
	low,heigh := 0,len(nums)-1
	for low <= heigh {
		middle := (low+heigh)/2
		if nums[middle] == target {
			return middle
		} else if nums[middle] > target {
			heigh = middle-1
		} else {
			low = middle+1
		}
	}
	return low
}*/
