package double_pointer

func removeElement(nums []int, val int) int {
	var slow int
	for _, v := range nums {
		if v == val {
			continue
		} else {
			nums[slow] = v
			slow++
		}
	}
	return slow
}

/*
func removeElement(nums []int, val int) int {
	slow := 0
    for _,v := range nums {
        if v != val {
            nums[slow] = v
            slow++
        } else {
            continue
        }
    }
    return slow
}
*/
