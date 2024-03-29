package hashmap

/*
给定两个字符串 s 和 p，找到 s 中所有 p 的 异位词 的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。

示例 1:
输入: s = "cbaebabacd", p = "abc"
输出: [0,6]
解释:
起始索引等于 0 的子串是 "cba", 它是 "abc" 的异位词。
起始索引等于 6 的子串是 "bac", 它是 "abc" 的异位词。
*/

/*
记录下p的长度，按照滑动窗口那样处理即可
*/

//func findAnagrams(s string, p string) []int {
//	l := len(p)
//	var cnt [26]int
//	for _, ch := range p {
//		cnt[ch-'a']++
//	}
//	for _, ch := range s {
//		if cnt[ch-'a'] == 0 {
//			continue
//		}
//	}
//}
