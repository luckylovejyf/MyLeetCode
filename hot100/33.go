package hot100

// 33. 搜索旋转排序数组 二分查找
// 转换后一定有一边是有序的
// 如果左边有序，且范围在左边，去左边查找，否则去右边
// 如果右边有序，且范围在右边，去右边查找，否则去左边

// 二分法总结
// 跳出条件 l<=r
// mid计算：减号计算
// r和l更新：mid-1，mid+1
// 查找相等的值返回mid
// 查找第一个大于等于的值，把<作为一类，返回l，判断是否越界
// 查找第一个小于等于的值，把>作为一类，返回r，判断是否越界
func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}

	l, r := 0, n-1
	for l <= r {
		mid := (r-l)/2 + l
		if nums[mid] == target {
			return mid
		}

		if nums[0] <= nums[mid] {
			if nums[0] <= target && nums[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}

		if nums[mid] <= nums[n-1] {
			if nums[mid] < target && nums[n-1] >= target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
