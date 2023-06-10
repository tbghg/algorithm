package stack_problem

type MinStack struct {
	minVal int //存储当前栈中的最小值
	s      []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (m *MinStack) Push(x int) {
	if len(m.s) == 0 {
		m.minVal = x
		m.s = append(m.s, 0)
		return
	}
	// 存储x与当前最小值的差值
	m.s = append(m.s, x-m.minVal)
	m.minVal = min(m.minVal, x)
}

func (m *MinStack) Pop() {
	m.minVal = min(m.Top(), m.minVal)
	m.s = m.s[:len(m.s)-1]
}

func (m *MinStack) Top() int {
	num := m.s[len(m.s)-1]
	if num < 0 {
		return m.minVal
	}
	return num + m.minVal
}

func (m *MinStack) Min() int {
	return m.minVal
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */

// [-2],[0],[-3]
//  -2   2   -1
