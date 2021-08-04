package week2

/*扫描方式
1 1 1
0 0 0
0 0 0
------
1 1 1
1 1 1
0 0 0
------
1 1 1
1 1 1
1 1 1
-------
0 0 0
1 1 1
0 0 0
-------
0 0 0
1 1 1
1 1 1
-------
…
自上往下，逐列遍历
*/
func numSubmatrixSumTarget(matrix [][]int, target int) int {
	ans := 0
	for i := range matrix {
		sum := make([]int, len(matrix[0]))
		for _, row := range matrix[i:] {
			for c, v := range row {
				//每列的前缀和，自上而下累加的值
				sum[c] += v
			}
			//枚举第一行的时候，1 0 0 ， 第一列的前缀和为1，当成一个独立的矩阵
			/*第二行时 	1 1 1
						1 0 0
			第一列
			1
			1 看成一个独立矩阵，然后向右拓展，最后扩散，把矩形中的所有子矩阵枚举完毕
			*/
			ans += subarrySum(sum, target)
		}
	}
	return ans
}
func subarrySum(nums []int, k int) (ans int) {
	mp := map[int]int{0: 1}
	for i, pre := 0, 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := mp[pre-k]; ok {
			ans += mp[pre-k]
		}
		mp[pre]++
	}
	return
}
