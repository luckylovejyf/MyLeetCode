package hot100

// 34. 在排序数组中查找元素的第一个和最后一个位置 二分查找
// start相当于找第一个大于等于的index，判断一下是否相同，如果不相同，说明没有target
// end相当于找第一个大于等于target+1的index，再减去1，只要start存在，end一定存在
func searchRange(nums []int, target int) []int {
	start := findEqualOrGreater(nums, target)
	if start >= len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	end := findEqualOrGreater(nums, target+1) - 1
	return []int{start, end}
}

func findEqualOrGreater(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (r-l)/2 + l
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
