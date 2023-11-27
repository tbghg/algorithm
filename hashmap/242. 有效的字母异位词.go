package hashmap

/*
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。

示例 1:
输入: s = "anagram", t = "nagaram"
输出: true
*/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mp := make(map[byte]int)
	for i := range s {
		mp[s[i]]++
	}
	for i := range t {
		mp[t[i]]--
		if mp[t[i]] < 0 {
			return false
		}
	}
	return true
}

func isAnagram1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var m = make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
		m[t[i]]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}
