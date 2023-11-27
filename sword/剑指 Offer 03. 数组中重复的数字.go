package sword

func findRepeatNumber(nums []int) int {
	var i int
	for i < len(nums) {
		if nums[i] == i {
			i++
			continue
		}
		// 一直交换，直到交换到正常位置或发现重复值
		// 发现重复值：要交换到值和当前值相等
		if nums[i] == nums[nums[i]] {
			return nums[i]
		}
		nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
	}
	return 0
}

//0 1 4 2 5 3 4
//0 1 5 2 4 3 4
//0 1 3 2 4 5 4
//0 1 2 3 4 5 4
