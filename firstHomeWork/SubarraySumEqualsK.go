package firstHomeWork

func subarraySum(nums []int, k int) int {
	/*先算前缀和
	  s1-n
	  s[r]-s[l]=k
	  s[l] = s[r]-k
	  如果提前统计前缀和数量，k==0时，会导致包含本身这个 因为val-val=0,但这却不是一个前缀和
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
