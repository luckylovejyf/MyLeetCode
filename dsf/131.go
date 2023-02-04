package dsf

// 113. 路径总和 II
// 深度优先遍历(当前、左节点、右节点)，先判断nil
// 加入当前节点，defer舍弃当前节点
// 计算剩余数据是否为0并且为叶节点
// 加入答案集，要新建slice，否则会被修改

func pathSum(root *TreeNode, targetSum int) [][]int {
	var ans [][]int
	var path []int
	var search func(*TreeNode, int)

	search = func(root *TreeNode, left int) {
		if root == nil {
			return
		}

		path = append(path, root.Val)
		defer func() { path = path[:len(path)-1] }()

		left -= root.Val
		if left == 0 && root.Left == nil && root.Right == nil {
			ans = append(ans, append([]int{}, path...))
			return
		}

		search(root.Left, left)
		search(root.Right, left)
	}

	search(root, targetSum)
	return ans
}

func pathSumAns(root *TreeNode, targetSum int) (ans [][]int) {
	var path []int
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, left int) {
		if node == nil {
			return
		}
		left -= node.Val
		path = append(path, node.Val)
		defer func() { path = path[:len(path)-1] }()
		if node.Left == nil && node.Right == nil && left == 0 {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		dfs(node.Left, left)
		dfs(node.Right, left)
	}
	dfs(root, targetSum)
	return
}
