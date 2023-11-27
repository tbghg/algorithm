package sword

func firstUniqChar1(s string) byte {
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}
	for _, ch := range s {
		if cnt[ch-'a'] == 1 {
			return byte(ch)
		}
	}
	return ' '
}

func firstUniqChar(s string) byte {
	n := len(s)
	cnt := [26]int{}
	for i := range cnt[:] {
		cnt[i] = n
	}
	for i, ch := range s {
		idx := ch - 'a'
		if cnt[idx] == n {
			cnt[idx] = i
		} else {
			cnt[idx] = n + 1
		}
	}
	for _, v := range cnt {
		if v < n {
			n = v
		}
	}
	if n == len(s) {
		return ' '
	}
	return s[n]
}
