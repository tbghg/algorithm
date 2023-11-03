package sword

func reverseLeftWords1(s string, n int) string {
	b := []byte(s)
	return string(b[n:]) + string(b[:n])
}

func reverseLeftWords(s string, n int) string {
	l := len(s)
	result := make([]byte, l)
	for i, ch := range s {
		result[(i+l-n)%l] = byte(ch)
	}
	return string(result)
}
