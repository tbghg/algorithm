package sword

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针
func reversePrint1(head *ListNode) (res []int) {
	p := head
	for p != nil {
		res = append(res, p.Val)
		p = p.Next
	}
	l, r := 0, len(res)-1
	for l < r {
		res[l], res[r] = res[r], res[l]
		l++
		r--
	}
	return
}

// 递归
func reversePrint(head *ListNode) []int {
	var res []int
	var f func(*ListNode)
	f = func(p *ListNode) {
		if p == nil {
			return
		}
		f(p.Next)
		res = append(res, p.Val)
	}
	f(head)
	return res
}
