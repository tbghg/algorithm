解决链表相关的题目，一般会用到哨兵/虚拟节点，如`dummy := &ListNode{Next: head}`

通过前驱节点进行访问，如`pre := dummy`，最后返回`dummy.Next`即可

Go语言中自带链表，里面函数挺多的，要用的时候再查就行，下面是对应的`example_test.go`

```go
package list_test

import (
	"container/list"
	"fmt"
)

func Example() {
	// Create a new list and put some numbers in it.
	l := list.New()
	e4 := l.PushBack(4)
	e1 := l.PushFront(1)
	l.InsertBefore(3, e4)
	l.InsertAfter(2, e1)

	// Iterate through list and print its contents.
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// Output:
	// 1
	// 2
	// 3
	// 4
}
```

值得一提的是go语言交换链表十分方便，大家可以对比一下下面两种写法

(写成一行，找到要换成的节点对应着写即可)

```go
package pointer_list
// 24.两两交换链表中的节点

type ListNode struct {
	Val  int
	Next *ListNode
}

// 注意for循环中交换节点的过程
func swapPairs1(head *ListNode) *ListNode {
	// 1 2 3 4 5 6 7 -> 2 1 4 3 6 5 7
	dummy := &ListNode{Next: head}
	pre, cur := dummy, head
	for cur != nil || cur.Next != nil {
		next := cur.Next
		pre.Next = next
		cur.Next = next.Next
		next.Next = cur
		pre = cur
		cur = cur.Next
	}
	return dummy.Next
}

// 三个节点，直接找到对应的Next，写成一行，找需要的赋值即可，简单方便
func swapPairs2(head *ListNode) *ListNode {
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
```
