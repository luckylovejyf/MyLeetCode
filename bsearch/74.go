package bsearch

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
