package dsf

// 错误
func findOrder(numCourses int, prerequisites [][]int) []int {
	m := make(map[int][]int)

	for _, prerequisite := range prerequisites {
		m[prerequisite[1]] = append(m[prerequisite[1]], prerequisite[0])
	}

	visited := make(map[int]bool)
	var path []int
	var dsf func(n int) ([]int, bool)
	dsf = func(n int) ([]int, bool) {
		if visited[n] {
			if len(path) == numCourses {
				return path, true
			} else {
				path = path[:len(path)-1]
				return nil, false
			}
		}

		visited[n] = true
		path = append(path, n)
		for _, i := range m[n] {
			ans, b := dsf(i)
			if b {
				return ans, true
			}
		}
		visited[n] = false
		path = path[:len(path)-1]
		return nil, false
	}

	for i, _ := range m {
		ans, b := dsf(i)
		if b {
			return ans
		}
	}

	return nil
}
