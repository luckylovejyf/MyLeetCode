package bsf

// 117. 填充每个节点的下一个右侧节点指针
// 广度优先，每一层通过nil区分

func connect117(root *Node) *Node {
	if root == nil {
		return nil
	}

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

		if q[i].Left != nil {
			q = append(q, q[i].Left)
		}
		if q[i].Right != nil {
			q = append(q, q[i].Right)
		}

		l = len(q)
	}

	for i := 0; i < len(q); i++ {
		if q[i] != nil {
			q[i].Next = q[i+1]
		}
	}

	return root
}
