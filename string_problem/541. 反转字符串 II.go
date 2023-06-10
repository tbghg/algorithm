package string_problem

/*
给定一个字符串 s 和一个整数 k，从字符串开头算起，每计数至 2k 个字符，就反转这 2k 字符中的前 k 个字符。
如果剩余字符少于 k 个，则将剩余字符全部反转。
如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。

示例 1：
输入：s = "abcdefg", k = 2
输出："bacdfeg"
*/

// 方法一
func reverseStr1(s string, k int) string {
	b := []byte(s)
	for i := 0; i < len(s); i += 2 * k {
		sub := b[i : i+min(k, len(s)-i)]
		for l, r := 0, len(sub)-1; l < r; l, r = l+1, r-1 {
			sub[l], sub[r] = sub[r], sub[l]
		}
	}
	return string(b)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 方法二
// 更加清晰易写，推荐
func reverseStr2(s string, k int) string {
	b := []byte(s)
	for i := 0; i < len(s); i += 2 * k {
		// 将每个长度为 k 的子串逆序
		if i+k <= len(s) {
			reverse(b[i : i+k])
		} else {
			reverse(b[i:])
		}
	}
	return string(b)
}

func reverse(sub []byte) {
	for l, r := 0, len(sub)-1; l < r; l, r = l+1, r-1 {
		sub[l], sub[r] = sub[r], sub[l]
	}
}
