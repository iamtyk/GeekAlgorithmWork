package week2

type ListNode struct {
	Val  int
	Next *ListNode
}

//朴素暴力循环
/*
func mergeKLists(lists []*ListNode) *ListNode {
	protect := &ListNode{}
	tail := protect
	for {
		isEmpty := true
		minIndex := 0
		minVale := math.MaxInt32
		//找最小
		for i := 0; i < len(lists); i++ {
			if lists[i] == nil {
				continue
			}
			isEmpty = false
			if minVale > lists[i].Val {
				minVale = lists[i].Val
				minIndex = i
			}
		}
		if isEmpty {
			break
		}
		temp := lists[minIndex]
		lists[minIndex] = lists[minIndex].Next
		tail.Next = temp
		tail = tail.Next
	}
	return protect.Next
}
*/

//分治，两两配对，合并在一起占用的空间和朴素一样都是5.2Mb
func mergeKLists(lists []*ListNode) *ListNode {
	return dfsMerge(lists, 0, len(lists)-1)
}

func dfsMerge(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	if l > r {
		return nil
	}

	mid := l + (r-l)/2
	return merge(dfsMerge(lists, l, mid), dfsMerge(lists, mid+1, r))
}

func merge(a, b *ListNode) *ListNode {
	protect := &ListNode{}
	tail := protect
	for a != nil || b != nil {
		if a == nil {
			tail.Next = b
			break
		}
		if b == nil {
			tail.Next = a
			break
		}
		var temp *ListNode
		if a.Val <= b.Val {
			temp = a
			a = a.Next
		} else {
			temp = b
			b = b.Next
		}
		tail.Next = temp
		tail = tail.Next
	}
	return protect.Next
}
