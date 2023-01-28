package dp

// timeout
func isInterleave(s1 string, s2 string, s3 string) bool {
	l1 := len(s1)
	l2 := len(s2)
	l3 := len(s3)

	if l1+l2 != l3 {
		return false
	}

	if l1 == 0 && l2 == 0 && l3 == 0 {
		return true
	}

	if l1 == 0 {
		return s2 == s3
	}
	if l2 == 0 {
		return s1 == s3
	}

	if s1[0] == s3[0] {
		if isInterleave(s1[1:], s2, s3[1:]) {
			return true
		}
	}

	if s2[0] == s3[0] {
		if isInterleave(s1, s2[1:], s3[1:]) {
			return true
		}
	}

	return false
}

// dp
func IsInterleaveDP(s1 string, s2 string, s3 string) bool {
	l1 := len(s1)
	l2 := len(s2)
	l3 := len(s3)

	if l1+l2 != l3 {
		return false
	}

	if l1 == 0 && l2 == 0 && l3 == 0 {
		return true
	}

	if l1 == 0 {
		return s2 == s3
	}
	if l2 == 0 {
		return s1 == s3
	}

	dp := make([][]bool, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]bool, l2+1)
	}

	dp[0][0] = true
	for i := 1; i <= l1; i++ {
		dp[i][0] = s1[:i] == s3[:i]
	}
	for i := 1; i <= l2; i++ {
		dp[0][i] = s2[:i] == s3[:i]
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1]) {
				dp[i][j] = true
			}
		}
	}

	return dp[l1][l2]
}
