package firstHomeWork

type ListNode struct {
	Val  int
	Next *ListNode
}

func PlusOne(digits []int) []int {
	career := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += career
		career = 0
		if digits[i] > 9 {
			digits[i] = 0
			career = 1
		}
	}

	var ans []int
	if career > 0 {
		ans = append(ans, career)
	}
	ans = append(ans, digits...)
	return ans
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//测试提交
	return &ListNode{}
}
