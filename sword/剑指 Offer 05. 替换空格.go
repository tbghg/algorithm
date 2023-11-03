package sword

func replaceSpace(s string) string {
	count := 0
	for _, ch := range s {
		if ch == ' ' {
			count++
		}
	}
	result := make([]byte, len(s)+count*2)
	var idx int
	for _, ch := range s {
		if ch == ' ' {
			result[idx] = '%'
			result[idx+1] = '2'
			result[idx+2] = '0'
			idx += 3
			continue
		}
		result[idx] = byte(ch)
		idx++
	}
	return string(result)
}
