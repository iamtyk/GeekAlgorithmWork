package week5

func minEatingSpeed(piles []int, H int) int {
	l := 1
	r := getmax(piles)
	for l <= r {
		mid := l + (r-l)/2
		if canfinish(piles, mid, H) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return r + 1
}

func getmax(piles []int) int {
	max := 0
	for _, v := range piles {
		if v > max {
			max = v
		}
	}
	return max
}
func canfinish(piles []int, speed, H int) bool {
	length := len(piles)
	time := 0
	for i := 0; i < length; i++ {
		t := piles[i] / speed
		if piles[i]%speed > 0 {
			t++
		}
		time = time + t
	}
	return time <= H
}
