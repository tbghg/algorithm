package stack_problem

// 两个栈来实现队列
// inStack、outStack
// 装入的时候直接放inStack
// 取出的时候，检查outStack里面有无数据
// 		有数据直接取出
// 		无数据则inStack逆着全部装里面，然后outStack取数据

type CQueue struct {
	inStack, outStack []int
}

func Constructor1() CQueue {
	return CQueue{}
}

func (c *CQueue) AppendTail(value int) {
	c.inStack = append(c.inStack, value)
}

func (c *CQueue) DeleteHead() int {
	if len(c.outStack) == 0 {
		if len(c.inStack) == 0 {
			return -1
		}
		c.in2out()
	}
	result := c.outStack[len(c.outStack)-1]
	c.outStack = c.outStack[:len(c.outStack)-1]
	return result
}

func (c *CQueue) in2out() {
	for len(c.inStack) > 0 {
		c.outStack = append(c.outStack, c.inStack[len(c.inStack)-1])
		c.inStack = c.inStack[:len(c.inStack)-1]
	}
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
