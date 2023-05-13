package pointer_list

/*
给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。
*/

// 思路：双指针，A+B和B+A的长度是一样的
//
//	如果有交点，相遇点为交点
//	如果没交点，二者遍历结束后同时为nil
//
// 时间复杂度：O(m+n)
// 空间复杂度：O(1)
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB
	for p1 != p2 {
		if p1.Next == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}
		if p2.Next == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1
}
