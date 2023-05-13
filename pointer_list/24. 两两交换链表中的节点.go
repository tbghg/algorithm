package pointer_list

/*
给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点
你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
*/

// 递归
func swapPairsRecursion(head *ListNode) *ListNode {
	// 1 2 3 4 5 6 7 -> 2 1 4 3 6 5 7
	if head == nil || head.Next == nil {
		return head
	}
	// 将目前的两位翻转，后面的节点直接递归
	result := swapPairs(head.Next.Next)
	next := head.Next
	head.Next, next.Next = result, head
	return next
}

// 迭代
func swapPairs(head *ListNode) *ListNode {
	// 1 2 3 4 5 6 7 -> 2 1 4 3 6 5 7
	dummy := &ListNode{Next: head}
	pre, cur := dummy, head
	for cur != nil && cur.Next != nil {
		next := cur.Next
		pre.Next, cur.Next, next.Next = next, next.Next, cur
		pre, cur = cur, cur.Next
	}
	return dummy.Next
}

// 方法三(gpt做的)
func swapPairsGPT(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for pre.Next != nil && pre.Next.Next != nil {
		pre = swap(pre)
	}
	return dummy.Next
}

func swap(pre *ListNode) *ListNode {
	a := pre.Next
	b := a.Next
	pre.Next, a.Next, b.Next = b, b.Next, a
	return a
}
