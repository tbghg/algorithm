package pointer_list

/*
题意：删除链表中等于给定值 val 的所有节点。
示例 1： 输入：head = [1,2,6,3,4,5,6], val = 6 输出：[1,2,3,4,5]
示例 2： 输入：head = [], val = 1 输出：[]
示例 3： 输入：head = [7,7,7,7], val = 7 输出：[]
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	// 1. 哨兵节点
	sentinel := &ListNode{Next: head}
	// 2. 前驱节点
	prev := sentinel
	// 3. 当前节点
	cur := head
	for cur != nil {
		if cur.Val == val {
			prev.Next = cur.Next
		} else {
			prev = cur
		}
		cur = cur.Next
	}
	return sentinel.Next
}
