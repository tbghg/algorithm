package pointer_list

/*
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]
*/

// 思路：快慢指针，快指针先走n步，然后快慢指针一起走，当快指针走到尾部时，慢指针指向的就是倒数第n个节点
//
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fast, slow := dummy, dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast, slow = fast.Next, slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

/*
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next:head}
    fast := dummy
    for i := 0 ; i < n; i++ {
        fast = fast.Next
    }
    slow := dummy
    for fast.Next != nil {
        fast = fast.Next
        slow = slow.Next
    }
    slow.Next = slow.Next.Next
    return dummy.Next
}

*/
