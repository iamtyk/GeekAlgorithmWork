package byteDanceFaceTry

import (
	"GeekAlgorithmWork/Global"
	"math"
	"strconv"
)

func Quick_sort(nums []int, l, r int) {
	if l+1 >= r {
		return
	}
	first, last := l, r-1
	key := nums[first]
	for first < last {
		for first < last && nums[last] >= key {
			last--
		}
		nums[first] = nums[last]
		for first < last && nums[first] <= key {
			first++
		}
		nums[last] = nums[first]
	}
	nums[first] = key
	Quick_sort(nums, l, first)
	Quick_sort(nums, first+1, r)
}

//左闭右开
func Merge_sort(nums, temp []int, l, r int) {
	if l+1 >= r {
		return
	}
	m := l + (r-l)/2
	Merge_sort(nums, temp, l, m)
	Merge_sort(nums, temp, m, r)
	p, i := l, l
	q := m
	for p < m || q < r {
		if q >= r || (p < m && nums[p] <= nums[q]) {
			temp[i] = nums[p]
			i++
			p++
		} else {
			temp[i] = nums[q]
			i++
			q++
		}
	}
	for i = l; i < r; i++ {
		nums[i] = temp[i]
	}

}

//最大子序和
func maxSubArray(nums []int) int {
	pre := make([]int, len(nums)+1)
	for i, val := range nums {
		pre[i+1] = pre[i] + val
	}
	preMin := make([]int, len(nums)+1)
	for i, _ := range nums {
		preMin[i+1] = Global.Min(preMin[i], pre[i+1])
	}
	ans := math.MinInt32
	for i := 1; i <= len(nums); i++ {
		//i之前的 -->j<=i-1
		ans = Global.Max(ans, pre[i]-preMin[i-1])
	}
	return ans
}

//差分
/*
	差分和前缀和是两个互逆运算
	差分数组，求前缀和就变成原数组
	前缀和数组，求差分变成原数组
	cal_b(cal_sum(A))=A
*/
/* 任何对于区间的操作，可以转化为两个关键点（事件）
事件的影响从l开始，在r+1处消失
累加影响得到答案
l +d    r+1   -d

  1  2  3  4  5
 10    -10
    20    -20
    25          (-25)
 10 45 -10 -20 0
 10 55 45  25 25
*/
//1109. 航班预订统计
func corpFlightBookings(bookings [][]int, n int) []int {
	delta := make([]int, n+2)
	for _, book := range bookings {
		fir := book[0]
		last := book[1]
		val := book[2]
		delta[fir] += val
		delta[last+1] -= val
	}
	sum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + delta[i]
	}
	return sum[1:]
}

func bigDataAdd(strA, strB string) (result string) {
	lenA := len(strA)
	lenB := len(strB)
	lenBase := 0
	if lenA > lenB {
		lenBase = lenA
	} else {
		lenBase = lenB
	}
	j := 0 //进位数
	result = ""
	inxA := lenA - 1
	inxB := lenB - 1
	itemA := 0
	itemB := 0
	for inxBase := lenBase - 1; inxBase >= 0; inxBase-- {
		itemA = 0
		if inxA >= 0 {
			itemA, _ = strconv.Atoi(string(strA[inxA]))
			inxA--
		}
		itemB = 0
		if inxB >= 0 {
			itemB, _ = strconv.Atoi(string(strB[inxB]))
			inxB--
		}
		sum := itemA + itemB + j
		if sum > 9 {
			j = 1
			sum = sum - 10
		} else {
			j = 0
		}
		result = strconv.Itoa(sum) + result
	}

	if j > 0 {
		result = strconv.Itoa(j) + result
	}

	return
}
