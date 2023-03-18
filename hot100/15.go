package hot100

import "sort"

// 15. 三数之和 双指针
// 先排序，固定第一个数字，第二三个数字从两边向中间靠
// 每次遍历的时候需要判断一下是否和之前一个相等
// 应为是有序的，所以随着第二个数字的递增，第三个数字是递减的
// 如果我们发现随着第一个元素的递增，第二个元素是递减的，那么就可以使用双指针的方法
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var ans [][]int

	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}

		third := n - 1
		for second := first + 1; second < n; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			for second < third && nums[second]+nums[third]+nums[first] > 0 {
				third--
			}

			if second == third {
				break
			}
			if nums[second]+nums[third]+nums[first] == 0 {
				ans = append(ans, []int{nums[second], nums[third], nums[first]})
			}
		}
	}
	return ans
}
