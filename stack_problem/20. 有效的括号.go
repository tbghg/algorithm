package stack_problem

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。

示例 1：
输入：s = "()"
输出：true
*/

func isValid(s string) bool {
	match := map[byte]byte{')': '(', ']': '[', '}': '{'}
	var stack []byte
	for _, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, byte(ch))
		} else {
			// 重点⚠️，如果stack为空，stack[-1]会报错
			// 需要先判断stack是否为空
			if len(stack) > 0 && stack[len(stack)-1] == match[byte(ch)] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
