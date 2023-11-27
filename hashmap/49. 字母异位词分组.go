package hashmap

import (
	"sort"
)

/*
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母通常恰好只用一次。

示例 1:
输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
*/

/*
方法一：异位词排序后都一样

由于互为字母异位词的两个字符串包含的字母相同，因此对两个字符串分别进行排序之后得到的字符串一定是相同的

故可以将排序之后的字符串作为哈希表的键
*/
func groupAnagrams1(strs []string) [][]string {
	mp := make(map[string][]string)
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		mp[string(s)] = append(mp[string(s)], str)
	}
	result := make([][]string, 0, len(mp))
	for _, v := range mp {
		result = append(result, v)
	}
	return result
}

/*
方法二：每个str对应一个[26]int，这个作为map的key
*/
func groupAnagrams(strs []string) [][]string {
	mp := make(map[[26]int][]string)
	for _, str := range strs {
		var cnt [26]int
		for _, ch := range str {
			cnt[ch-'a']++
		}
		mp[cnt] = append(mp[cnt], str)
	}
	result := make([][]string, 0, len(mp))
	for _, v := range mp {
		result = append(result, v)
	}
	return result
}
