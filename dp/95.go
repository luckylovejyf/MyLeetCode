package dp

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n <= 0 {
		return []*TreeNode{nil}
	}

	return helper(1, n)
}

func helper(start int, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	res := make([]*TreeNode, 0)
	for i := start; i <= end; i++ {
		left := helper(start, i-1)
		right := helper(i+1, end)

		for _, leftNode := range left {
			for _, rightNode := range right {
				res = append(res, &TreeNode{i, leftNode, rightNode})
			}
		}
	}

	return res
}
