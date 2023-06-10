package double_pointer

func backspaceCompare(s string, t string) bool {
	//fmt.Println(getResult(s))
	//fmt.Println(getResult(t))
	return getResult(s) == getResult(t)
}

func getResult(s string) string {
	var slow int
	b := []byte(s)
	for _, ch := range b {
		if ch == '#' {
			if slow > 0 {
				slow--
			}
		} else {
			b[slow] = ch
			slow++
		}
	}
	return string(b[:slow])
}

// 官方答案
//
// 一个字符是否会被删掉，只取决于该字符后面的退格符，而与该字符前面的退格符无关。
// 因此当我们逆序地遍历字符串，就可以立即确定当前字符是否会被删掉。
func backspaceCompare1(s, t string) bool {
	skipS, skipT := 0, 0
	i, j := len(s)-1, len(t)-1
	for i >= 0 || j >= 0 {
		for i >= 0 {
			if s[i] == '#' {
				skipS++
				i--
			} else if skipS > 0 {
				skipS--
				i--
			} else {
				break
			}
		}
		for j >= 0 {
			if t[j] == '#' {
				skipT++
				j--
			} else if skipT > 0 {
				skipT--
				j--
			} else {
				break
			}
		}
		if i >= 0 && j >= 0 {
			if s[i] != t[j] {
				return false
			}
		} else if i >= 0 || j >= 0 {
			return false
		}
		i--
		j--
	}
	return true
}
