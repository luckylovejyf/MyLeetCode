package dsf

// 222. 完全二叉树的节点个数
// 深度优先、广度优先

func countNodes(root *TreeNode) int {
	var ans int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		ans++
		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	return ans
}
