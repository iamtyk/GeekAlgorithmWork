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

//不能直接使用优美子的查找
func numberOfSubarrays(nums []int, k int) int {
	count := map[int]int{}
	s := make([]int, len(nums)+1)
	//计算前缀和
	for i, val := range nums {
		s[i+1] = s[i] + val
	}
	//计数
	for _, val := range s {
		count[val]++
	}
	//查找
	ans := 0
	for i := 1; i < len(s); i++ {
		val := s[i]
		//这里就导致无序了，不能保证val>=k,遍历的一定是做左边，因为可能为负数、0
		ans += count[val-k]
	}

	return ans
}
