package week3

func solve(board [][]byte) {
	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}
	visist := make([][]bool, len(board))
	for i := range visist {
		visist[i] = make([]bool, len(board[0]))
	}
	m, n := len(board), len(board[0])
	//bfs
	q := [][]int{}
	ans := [][][]int{}
	for row, arr := range board {
		for col, ch := range arr {
			if visist[row][col] {
				continue
			}
			visist[row][col] = true
			if ch == 'X' {
				continue
			}
			q = append(q, []int{row, col})
			inserO := [][]int{{0}} //标记插入的O是否临近边缘，0表示没有
			for len(q) > 0 {
				p := q[0]

				x, y := p[0], p[1]
				inserO = append(inserO, []int{x, y})
				if x == 0 || x == m-1 || y == 0 || y == n-1 {
					//是边缘，置标记
					inserO[0][0] = 1
				}
				q = q[1:]
				for i := 0; i < 4; i++ {
					nx := x + dx[i]
					ny := y + dy[i]
					if nx < 0 || ny < 0 || nx >= m || ny >= n {
						continue
					}

					if visist[nx][ny] {
						continue
					}

					visist[nx][ny] = true
					if board[nx][ny] == 'X' {
						continue
					}
					q = append(q, []int{nx, ny})
				}
			}
			//fmt.Println(inserO)
			//找完一个团子，加入结果
			ans = append(ans, inserO)
		}
	} //bfs结束

	//填充X
	for _, arr := range ans {
		if arr[0][0] == 1 {
			continue
		}
		for i := 1; i < len(arr); i++ {
			board[arr[i][0]][arr[i][1]] = 'X'
		}
	}

}
