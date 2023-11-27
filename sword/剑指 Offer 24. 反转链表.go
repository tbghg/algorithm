package sword

// 迭代
func reverseList1(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre, cur = cur, next
	}
	return pre
}

// 递归
// 返回的需要是最后的
// 到最后直接返回node，每一层都把这个node直接返回
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		// 返回最后的节点
		return head
	}
	// 保存下来节点，为了return
	newHead := reverseList(head.Next)
	// 反转
	head.Next.Next = head
	head.Next = nil
	return newHead
}
