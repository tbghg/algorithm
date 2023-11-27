package hashmap

/*
给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。
如果可以，返回 true ；否则返回 false 。
magazine 中的每个字符只能在 ransomNote 中使用一次。

示例 1：

输入：ransomNote = "a", magazine = "b"
输出：false
*/

func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[byte]int)
	for i := 0; i < len(magazine); i++ {
		m[magazine[i]]++
	}
	for i := 0; i < len(ransomNote); i++ {
		if m[ransomNote[i]] == 0 {
			return false
		} else {
			m[ransomNote[i]]--
		}
	}
	return true
}

func canConstruct1(ransomNote string, magazine string) bool {
	var mp [26]int
	for _, ch := range magazine {
		mp[ch-'a']++
	}
	for _, ch := range ransomNote {
		mp[ch-'a']--
		if mp[ch-'a'] < 0 {
			return false
		}
	}
	return true
}
