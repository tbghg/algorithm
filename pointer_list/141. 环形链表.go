package pointer_list

func hasCycle(head *ListNode) bool {
	f, s := head, head
	for f != nil || f.Next != nil {
		f = f.Next.Next
		s = s.Next
		if f == s {
			return true
		}
	}
	return false
}
