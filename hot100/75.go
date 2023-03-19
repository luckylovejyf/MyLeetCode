package hot100

// 75. 颜色分类 多指针
// p0 p1指针分别记录0和1的下一个位置 p1>=p0
// 从头开始向后移动，遇到0，与p0交换；如果p0<p1，说明将1交换出去了，需要与p1交换回来 之后p0和p1都需要++
// 遇到1，与p1交换，p1++

func sortColors(nums []int) {
	n := len(nums)
	if n == 0 {
		return
	}

	p0, p1 := 0, 0
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			nums[p0], nums[i] = nums[i], nums[p0]
			if p0 < p1 {
				nums[p1], nums[i] = nums[i], nums[p1]
			}
			p0++
			p1++
		} else if nums[i] == 1 {
			nums[p1], nums[i] = nums[i], nums[p1]
			p1++
		}
	}
}
