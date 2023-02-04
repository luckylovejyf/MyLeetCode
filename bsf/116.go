package bsf

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 116. 填充每个节点的下一个右侧节点指针
// 广度优先，每一层通过nil区分

func connect(root *Node) *Node {
	var q []*Node

	q = append(q, root)
	q = append(q, nil)

	l := len(q)
	for i := 0; i < l; i++ {
		if q[i] == nil && i == len(q)-1 {
			break
		}
		if q[i] == nil {
			q = append(q, nil)
			continue
		}

		q = append(q, q[i].Left)
		q = append(q, q[i].Right)

		l = len(q)
	}

	for i := 0; i < len(q); i++ {
		if q[i] != nil {
			q[i].Next = q[i+1]
		}
	}

	return root
}
