package test

import (
	"fmt"
	"sort"
)

func Solution(A []int) int {
	sort.Ints(A)
	m := 1
	for _, i := range A {
		if i <= 0 {
			continue
		}
		if i > m {
			break
		}
		if i == m {
			m++
		}
	}
	return m
}

func Solution1(N int) {
	var enable_print int
	enable_print = N % 10
	for N > 0 {
		if enable_print == 0 && N%10 != 0 {
			enable_print = 1
			fmt.Print(N % 10)
		} else if enable_print == 1 {
			fmt.Print(N % 10)
		}
		N = N / 10
	}
}

func Solution3(A []int) int {
	if len(A) == 0 {
		return 0
	}

	l := len(A)
	var ans int
	isAllZero := true
	for _, n := range A {
		if n != 0 {
			isAllZero = false
			break
		}
	}
	if isAllZero {
		ans = l * (l + 1) / 2
		if ans > 1000000000 {
			return -1
		} else {
			return ans
		}
	}

	dp := make([][]int, l)
	for i := 0; i < l; i++ {
		dp[i] = make([]int, l)
	}

	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			if i == j {
				dp[i][j] = A[i]
			} else {
				dp[i][j] = dp[i][j-1] + A[j]
			}
		}
	}

	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			if dp[i][j] == 0 {
				ans++
			}
		}
	}

	if ans > 1000000000 {
		return -1
	} else {
		return ans
	}
}

func Solution2(A []int, B []int, N int) int {
	ms := make(map[int]int, N)
	m := len(A)
	for i := 0; i < m; i++ {
		ms[A[i]]++
		ms[B[i]]++
	}

	var ans int
	for i := 0; i < m; i++ {
		cur := ms[A[i]] + ms[B[i]] - 1
		ans = max(ans, cur)
	}
	return ans
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
