package hot100

// 19. 删除链表的倒数第 N 个结点 双指针
// 第一个指针先向前移动n，然后两个指针同时移动
// 注意两个指针之间的间隔应该是n+1
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dumpHead := &ListNode{
		Val:  0,
		Next: head,
	}

	first := head
	for i := 0; i < n; i++ {
		first = first.Next
	}

	second := dumpHead
	for true {
		if first == nil {
			break
		}
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return dumpHead.Next
}
