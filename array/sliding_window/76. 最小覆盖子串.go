package sliding_window

import "math"

func minWindow(s string, t string) string {
	mp := make(map[byte]int)
	lenT := len(t)
	for _, ch := range t {
		mp[byte(ch)]++
	}
	var l, result int
	length := math.MaxInt32
	for r, ch := range s {
		if mp[byte(ch)] > 0 {
			lenT--
		}
		mp[byte(ch)]--
		// 判断是否已满足条件
		if lenT <= 0 {
			// 已满足条件，for循环左侧移动，缩小长度
			for {
				// 非t中的字符，右侧移动时肯定把它包含了，所以不可能为0
				// t中的字符如果为0，那现在这个字符的数量正好，把这个抛了就不满足了，所以break
				if mp[s[l]] == 0 {
					break
				}
				mp[s[l]]++
				l++
			}
			// 和原本的答案比较，看看是否要替换
			if length > r-l+1 {
				result = l
				length = r - l + 1
			}
		}
	}
	if length == math.MaxInt32 {
		return ""
	}
	return s[result : result+length]
}

/*
func minWindow(s string, t string) string {
	mp := make(map[byte]int)
	lenT := len(t)
	for _, ch := range t {
		mp[byte(ch)]++
	}
	var begin, result int
	length := math.MaxInt
	for end, ch := range s {
		if mp[byte(ch)] > 0 {
			lenT--
		}
		mp[byte(ch)]--
		if lenT <= 0 {
			for {
				if mp[s[begin]] == 0 {
					break
				} else {
					mp[s[begin]]++
					begin++
				}
			}
			if (end - begin + 1) < length {
				length = end - begin + 1
				result = begin
			}
		}
	}
	if length == math.MaxInt {
		return ""
	}
	return s[result : result+length]
}
*/
