package firstHomeWork

func subarraySum(nums []int, k int) int {
	/*先算前缀和
	  s1-n
	  s[r]-s[l]=k
	  s[l] = s[r]-k
	*/
	countMap := map[int]int{}
	countMap[0] = 1
	totalVal, ans := 0, 0
	for _, val := range nums {
		totalVal += val
		if _, ok := countMap[totalVal-k]; ok {
			ans += countMap[totalVal-k]
		}
		countMap[totalVal] += 1
	}

	return ans
}
