package hot100

// 31. 下一个排列
// 下一个排列有两种情况：1. 当前不是最大值，下一个排列一定比当前大 2. 当前是最大值；先讨论第一种情况
// 下一个比当前大，即通过将一个右边的大数和左边的小数交换一下
// 要变大的幅度尽可能小，所以要尽可能靠右（低位）去找一个较小的数字：从后遍历，找到第一个正序的，就是我们要找的最靠右的小数
// 要变大的幅度尽可能小，所以要在右边找一个最小的数字作为大数：从后遍历，找到第一个大于小数的
// 要变大的幅度尽可能小，所以交换完成之后，保证大数后面是正序，因为大数的位置是逆序的（跟我们找位置的算法有关），所以reverse一下就可以
// 第二种情况，找第一个小数会找到-1，此时直接reverse一下就可以
func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	for ; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}

	if i >= 0 {
		j := n - 1
		for ; j > i; j-- {
			if nums[j] > nums[i] {
				break
			}
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	reverse(nums[i+1:])

}

func reverse(nums []int) {
	i, j := 0, len(nums)-1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
