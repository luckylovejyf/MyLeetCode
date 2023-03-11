package hot100

type ListNode struct {
	Val  int
	Next *ListNode
}

// 2. 两数相加
// 从头开始遍历，保存每两个对应节点相加后的结果，判断是否进位
// 若其中一个遍历提前结束，则单独遍历剩余节点
// 最后判断进位是否有值，需要额外添加节点
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	cur := head
	carry := 0
	for l1 != nil && l2 != nil {
		n1, n2 := l1.Val, l2.Val
		sum := (n1 + n2 + carry) % 10
		carry = (n1 + n2 + carry) / 10
		tmp := &ListNode{Val: sum}
		cur.Next = tmp
		cur = tmp
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		n1 := l1.Val
		sum := (n1 + carry) % 10
		carry = (n1 + carry) / 10
		tmp := &ListNode{Val: sum}
		cur.Next = tmp
		cur = tmp
		l1 = l1.Next
	}

	for l2 != nil {
		n1 := l2.Val
		sum := (n1 + carry) % 10
		carry = (n1 + carry) / 10
		tmp := &ListNode{Val: sum}
		cur.Next = tmp
		cur = tmp
		l2 = l2.Next
	}

	if carry > 0 {
		tmp := &ListNode{Val: carry}
		cur.Next = tmp
		cur = tmp
	}

	return head.Next
}
