package hot100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 114. 二叉树展开为链表 前序遍历
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	var nodes []*TreeNode
	var preTravel func(root *TreeNode)
	preTravel = func(root *TreeNode) {
		if root == nil {
			return
		}

		nodes = append(nodes, root)
		preTravel(root.Left)
		preTravel(root.Right)
	}
	preTravel(root)
	if len(nodes) <= 1 {
		return
	}
	for i := 1; i < len(nodes); i++ {
		nodes[i-1].Left = nil
		nodes[i-1].Right = nodes[i]
	}
}
