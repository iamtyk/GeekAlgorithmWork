package week4

import "container/heap"

type myheap struct {
	dataSlic []int
	big      bool
}

func (h *myheap) Less(i, j int) bool {
	if h.big {
		return (*h).dataSlic[i] > (*h).dataSlic[j]
	} else {
		return (*h).dataSlic[i] < (*h).dataSlic[j]
	}
}

func (h *myheap) Swap(i, j int) {
	(*h).dataSlic[i], (*h).dataSlic[j] = (*h).dataSlic[j], (*h).dataSlic[i]
}

func (h *myheap) Len() int {
	return len((*h).dataSlic)
}

func (h *myheap) Pop() (v interface{}) {
	(*h).dataSlic, v = (*h).dataSlic[:h.Len()-1], (*h).dataSlic[h.Len()-1]
	return
}

func (h *myheap) Push(v interface{}) {
	(*h).dataSlic = append((*h).dataSlic, v.(int))
}

type MedianFinder struct {
	maxHeap *myheap
	minHeap *myheap
}

/** initialize your data structure here. */
func Constructorlc295() MedianFinder {
	return MedianFinder{
		maxHeap: &myheap{big: true}, //大根堆，存储更小的一半
		minHeap: new(myheap),        //小根堆，存储大的一半
	}
}

func (t *MedianFinder) AddNum(num int) {
	if t.maxHeap.Len() == 0 {
		t.maxHeap.Push(num)
		return
	}
	if num <= t.maxHeap.dataSlic[0] {
		heap.Push(t.maxHeap, num)
	} else {
		heap.Push(t.minHeap, num)
	}

	//两堆的个数差距不能超过1
	if t.maxHeap.Len() > t.minHeap.Len()+1 {
		heap.Push(t.minHeap, heap.Pop(t.maxHeap))
	}

	//大根堆默认要比小根堆多，否则就是相等，这样奇数个时只要去大根堆取就可以了
	if t.maxHeap.Len() < t.minHeap.Len() {
		heap.Push(t.maxHeap, heap.Pop(t.minHeap))
	}
}

func (t *MedianFinder) FindMedian() float64 {
	if t.maxHeap.Len() > t.minHeap.Len() {
		return float64(t.maxHeap.dataSlic[0])
	} else {
		return float64(t.maxHeap.dataSlic[0]+t.minHeap.dataSlic[0]) / 2
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
