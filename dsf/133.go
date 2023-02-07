package dsf

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	m := make(map[int]*Node)

	var dsf func(cur *Node)
	dsf = func(cur *Node) {
		if node == nil {
			return
		}
		_, ok := m[cur.Val]
		if ok {
			return
		}

		newN := new(Node)
		newN.Val = cur.Val
		m[cur.Val] = newN
		for _, neighbor := range cur.Neighbors {
			n, ok := m[neighbor.Val]
			if ok {
				n.Neighbors = append(n.Neighbors, newN)
				newN.Neighbors = append(newN.Neighbors, n)
			}
			dsf(neighbor)
		}

	}

	ans := new(Node)
	ans.Val = node.Val
	m[ans.Val] = ans
	for _, neighbor := range node.Neighbors {
		dsf(neighbor)
	}

	return ans
}

// 133. 克隆图
// 需要记录是否克隆过该节点的map
// 深度优先搜索，如果为nil或者克隆过，直接返回
// 创建克隆节点，val复制
// 遍历原节点neighbor，克隆neighbor并赋值

func cloneGraphAns(node *Node) *Node {
	if node == nil {
		return node
	}

	m := make(map[int]*Node)

	var dsf func(cur *Node) *Node
	dsf = func(cur *Node) *Node {
		if node == nil {
			return node
		}
		if n, ok := m[cur.Val]; ok {
			return n
		}

		cloneN := new(Node)
		cloneN.Val = cur.Val
		m[cloneN.Val] = cloneN
		for _, neighbor := range cur.Neighbors {
			cloneN.Neighbors = append(cloneN.Neighbors, dsf(neighbor))
		}
		return cloneN
	}

	return dsf(node)
}
