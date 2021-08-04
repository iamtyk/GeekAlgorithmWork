package week3

func findOrder(numCourses int, prerequisites [][]int) []int {
	edges := make([][]int, numCourses)
	inDeg := make([]int, numCourses)
	ans := []int{}
	//加边
	for _, arr := range prerequisites {
		edges[arr[1]] = append(edges[arr[1]], arr[0])
		inDeg[arr[0]]++
	}

	var bfs func()
	bfs = func() {
		//从所有入度为0的边出发
		q := []int{}
		for i, val := range inDeg {
			if val == 0 {
				q = append(q, i)
			}
		}
		for len(q) > 0 {
			ans = append(ans, q[0])
			x := q[0]
			q = q[1:]
			for _, y := range edges[x] {
				inDeg[y]--
				if inDeg[y] == 0 {
					q = append(q, y)
				}
			}
		}
	}
	bfs()
	if len(ans) < numCourses {
		return []int{}
	}
	return ans
}
