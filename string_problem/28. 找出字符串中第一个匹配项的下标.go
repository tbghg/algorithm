package string_problem

import "strings"

/*
给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）
如果 needle 不是 haystack 的一部分，则返回 -1

输入：haystack = "sadbutsad", needle = "sad"
输出：0
解释："sad" 在下标 0 和 6 处匹配。
第一个匹配项的下标是 0 ，所以返回 0 。
*/

// 方法一：直接调用strings.Index
//
// 面试官：回去等电话吧
func strStr1(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

// 方法二：KMP
func strStr(haystack string, needle string) int {
	return 0
}

/*
func getNext(needle []byte) []int {
	i, j := 0, -1
	next := make([]int, len(needle)+1)
	next[0] = -1
	for i < len(needle) {
		if j == -1 || needle[i] == needle[j] {
			i++
			j++
			next[i] = j
		}else{
			j = next[j]
		}
	}
	return next
}

func strStr(haystack string, needle string) int {
	byteNeedle := []byte(needle)
	byteHayStack := []byte(haystack)
	next := getNext(byteNeedle)
	i, j := 0, 0
	for i < len(haystack) && j < len(needle) {
		// fmt.Println(i,j)
		if j == -1 || byteNeedle[j] == byteHayStack[i] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == len(needle) {
		return i - j
	}
	return -1
}
*/
