package dp

func Rob(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	if l == 1 {
		return nums[0]
	}

	if l == 2 {
		return max(nums[0], nums[1])
	}

	dpSteal := make([]int, l)
	dpNotSteal := make([]int, l)
	dpSteal[0] = nums[0]
	dpNotSteal[0] = 0
	for i := 1; i < l-1; i++ {
		dpSteal[i] = dpNotSteal[i-1] + nums[i]
		dpNotSteal[i] = max(dpSteal[i-1], dpNotSteal[i-1])
	}
	first := max(dpSteal[l-2], dpNotSteal[l-2])
	dpSteal[1] = nums[1]
	dpNotSteal[1] = 0
	for i := 2; i < l; i++ {
		dpSteal[i] = dpNotSteal[i-1] + nums[i]
		dpNotSteal[i] = max(dpSteal[i-1], dpNotSteal[i-1])
	}
	second := max(dpSteal[l-1], dpNotSteal[l-1])

	return max(first, second)
}
