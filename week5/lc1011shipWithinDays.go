package week5

func shipWithinDays(weights []int, D int) int {
	left, right, mid := 0, 0, 0
	for _, v := range weights {
		right += v
		if v > left {
			left = v
		}
	}
	for left < right {
		mid = (left + right) / 2
		if canFinish(weights, mid, D) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// 指定运载能力cap能否在D天内运完所有货物
func canFinish(weights []int, cap, D int) bool {
	daySum, day := 0, 1
	for i := 0; i < len(weights); {
		if daySum+weights[i] <= cap {
			daySum += weights[i]
			i++
		} else {
			daySum, day = 0, day+1
		}
	}
	return day <= D
}
