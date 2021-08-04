package week2

func robotSim(commands []int, obstacles [][]int) int {
	obsMap := map[int64]int{}
	maxDis := 0
	for _, val := range obstacles {
		obsMap[calHash(val[0], val[1])] = 1
	}
	yDir := []int{1, 0, -1, 0}
	xDir := []int{0, 1, 0, -1}
	index := 0
	ans := []int{0, 0}
	for i := 0; i < len(commands); i++ {
		if commands[i] == -1 {
			index += 1
			if index == 4 {
				index = 0
			}
			continue
		}
		if commands[i] == -2 {
			index -= 1
			if index == -1 {
				index = 3
			}
			continue
		}
		for j := 1; j <= commands[i]; j++ {
			tx := ans[0] + 1*xDir[index]
			ty := ans[1] + 1*yDir[index]
			if _, ok := obsMap[calHash(tx, ty)]; ok {
				break
			}
			ans[0], ans[1] = tx, ty
		}
		if ans[0]*ans[0]+ans[1]*ans[1] > maxDis {
			maxDis = ans[0]*ans[0] + ans[1]*ans[1]
		}
	}
	return maxDis
}
func calHash(x, y int) int64 {
	return (int64(x)+30000)*600000 + int64(y) + 30000
}
