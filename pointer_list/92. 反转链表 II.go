package pointer_list

/*
给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right
请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表

输入：head = [1,2,3,4,5], left = 2, right = 4
输出：[1,4,3,2,5]

输入：head = [5], left = 1, right = 1
输出：[5]

其中：
	+ 链表中节点数目为 n
	+ 1 <= n <= 500
	+ -500 <= Node.val <= 500
	+ 1 <= left <= right <= n
*/

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	// 找到要开始交换的节点的前一个，放在pre
	for i := 1; i < left; i++ {
		pre = pre.Next
	}
	// cur为开始要交换的节点
	cur := pre.Next
	// 1 2 3 4 5 6 == left=2 right=5
	// 第一轮：1 | 3 2 4 5 | 6   3放到1前
	// 第二轮：1 | 4 3 2 5 | 6   4放到1前
	// 第三轮：1 | 5 4 3 2 | 6   5放到1前
	// pre始终指向1，cur始终指向2，cur.Next为需要放到1前的那个
	for i := left; i < right; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummy.Next
}

/*func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	llNode := dummy
	for i := 1; i < left; i++ {
		llNode = llNode.Next
	}
	lNode := llNode.Next
	pre := lNode
	cur := pre.Next
	flag := right - left
	for flag > 0 {
		next := cur.Next
		cur.Next = pre
		pre, cur = cur, next
		flag--
	}
	llNode.Next = pre
	lNode.Next = cur
	return dummy.Next
}
*/
