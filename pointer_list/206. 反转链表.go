package pointer_list

/*
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
*/

// 双指针法
func reverseListDoublePoint(head *ListNode) *ListNode {
	// 初始的head应该指向nil
	// pre,cur,next
	var pre, cur *ListNode
	cur = head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre, cur = cur, next
	}
	return pre
}

// 递归
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// return返回的值是原本最末端的值，假定为result
	// result是最后head == nil || head.Next == nil那里返回的
	result := reverseList(head.Next)
	// 反转
	head.Next.Next = head
	// 当前的head.Next置空(实现最后的指针指向nil)
	head.Next = nil
	return result
}

/*
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    result := reverseList(head.Next)
    head.Next.Next = head
    head.Next = nil
    return result
}
*/
