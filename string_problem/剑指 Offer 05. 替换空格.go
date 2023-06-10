package string_problem

import "strings"

/*
请实现一个函数，把字符串 s 中的每个空格替换成"%20"

示例 1：
	输入：s = "We are happy."
	输出："We%20are%20happy."
限制：
	0 <= s 的长度 <= 10000
*/

// 方法一：直接调用
//
// 会被面试官打
func replaceSpace1(s string) string {
	return strings.Replace(s, " ", "%20", -1)
}
