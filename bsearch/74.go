package bsearch

// 74. 搜索二维矩阵
// 首先二分查找第一列，找到<=target的一行，然后二分查找这一行
// 二分查找<=：l=-1, r=len-1; mid=(r-l+1)/2+l; l=mid, r= mid-1 最后要判断l<0
// 二分查找==；l=0, r=len-1; mid=(r-l)/2+l; l=mid+1, r=mid-1

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	l, r := -1, len(matrix)-1
	for l < r {
		mid := (r-l+1)/2 + l

		if matrix[mid][0] <= target {
			l = mid
		} else {
			r = mid - 1
		}
	}

	index := l
	if index < 0 {
		return false
	}
	l, r = 0, len(matrix[index])-1
	for l <= r {
		mid := (r-l)/2 + l
		if matrix[index][mid] == target {
			return true
		}
		if matrix[index][mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}
