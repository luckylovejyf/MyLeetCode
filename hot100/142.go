package hot100

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 142. 环形链表 II 双指针
// fast走两步 slow走一步
// fast == nil 没有环
// fast与slow相遇，有环；假设开始到环的起点距离为a，环的起点到相遇点为b，相遇点到环的起点为c，相遇后fast走了n圈环
// fast走过的距离 = a+n(b+c)+b
// slow走过的距离 = a+b
// fast = 2 * slow
// fast = 2 * slow
// a =c+(n-1)(b+c) 从相遇点到环的起点+n-1圈环，刚好等于起点到环的起点的距离
// 所以相遇后，slow和third指针同时出发，刚好在环的起点相遇
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	fast, slow := head, head
	for true {
		if fast == nil || fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			third := head
			for true {
				if slow == third {
					return third
				}
				slow = slow.Next
				third = third.Next
			}
		}
	}
	return nil
}
