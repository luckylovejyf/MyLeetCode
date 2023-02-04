package dsf

// 99. 恢复二叉搜索树
// 中序遍历：递归左节点，处理本节点，递归右节点
// 中序遍历应该是递增序列，交换的两个会破坏递增结果
// 交换的两个不相邻、交换的两个相邻
// 在树中找到对应的位置，替换即可

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	nums := inOrder(root)

	x, y := find2(nums)

	recover2(root, 2, x, y)
}

func recover2(root *TreeNode, count int, x int, y int) {
	if root == nil {
		return
	}

	if root.Val == x {
		root.Val = y
		count--
		if count == 0 {
			return
		}
	} else if root.Val == y {
		root.Val = x
		count--
		if count == 0 {
			return
		}
	}

	recover2(root.Left, count, x, y)
	recover2(root.Right, count, x, y)
}

func find2(nums []int) (int, int) {
	index1, index2 := -1, -1

	for i := 0; i < len(nums); i++ {
		if i < len(nums)-1 && nums[i] > nums[i+1] {
			if index1 == -1 {
				index1 = i
			} else {
				index2 = i + 1
				break
			}
		}
	}

	if index2 == -1 {
		index2 = index1 + 1
	}

	return nums[index1], nums[index2]
}

func inOrder(root *TreeNode) []int {
	nums := make([]int, 0)
	if root == nil {
		return nums
	}

	nums = append(nums, inOrder(root.Left)...)
	nums = append(nums, root.Val)
	nums = append(nums, inOrder(root.Right)...)

	return nums
}
