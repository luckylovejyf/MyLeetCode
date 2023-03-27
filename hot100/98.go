package hot100

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 98. 验证二叉搜索树 中序遍历
// 二叉搜索树中序遍历后递增
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var vals []int
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		inorder(node.Left)
		vals = append(vals, node.Val)
		inorder(node.Right)
	}

	inorder(root)

	if len(vals) == 0 || len(vals) == 1 {
		return true
	}

	for i := 1; i < len(vals); i++ {
		if vals[i] <= vals[i-1] {
			return false
		}
	}
	return true
}
