package bsearch

// 33. 搜索旋转排序数组
// 二分查找，如果右面是增序且在右面范围内，在右边寻找，否则去左边寻找；如果左边是增序且在左边范围内，在左边寻找，否则去右边寻找

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

		if nums[mid] <= nums[n-1] {
			if nums[mid] < target && target <= nums[n-1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}

		if nums[0] <= nums[mid] {
			if nums[0] <= target && target <= nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}

	return -1
}
