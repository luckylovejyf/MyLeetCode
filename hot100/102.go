package hot100

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 102. 二叉树的层序遍历 层序遍历
// 先将根节点入队，然后塞入换行标识nil
// 出队，如果是nil并且队列为空，结束
// 如果是nil并且队列不为空，要换行，入队换行标识nil
// 如果不是nil，记录值，将左右子节点不是nil的入队
func levelOrder(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}

	var q []*TreeNode
	q = append(q, root)
	q = append(q, nil)

	var curLevel []int

	for true {
		cur := q[0]
		q = q[1:]
		if cur == nil && len(q) == 0 {
			ans = append(ans, append([]int{}, curLevel...))
			break
		}
		if cur == nil {
			q = append(q, nil)
			ans = append(ans, append([]int{}, curLevel...))
			curLevel = []int{}
			continue
		}

		curLevel = append(curLevel, cur.Val)
		if cur.Left != nil {
			q = append(q, cur.Left)
		}
		if cur.Right != nil {
			q = append(q, cur.Right)
		}
	}
	return ans
}
