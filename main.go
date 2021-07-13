package main

import (
	"GeekAlgorithmWork/Global"
	"GeekAlgorithmWork/byteDanceFaceTry"
	"container/heap"
	"fmt"
	"sort"
)

//大根堆
type myHeap [][]int

func (h *myHeap) Less(i, j int) bool {
	return (*h)[i][0] > (*h)[j][0]
}

func (h *myHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *myHeap) Len() int {
	return len(*h)
}

func (h *myHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *myHeap) Push(v interface{}) {
	*h = append(*h, v.([]int))
}
func main() {
	h := new(myHeap)
	h.Push(1)
	testMap := map[int]bool{}
	fmt.Println(testMap[1])
	testArrStr := "abc"
	strSin := testArrStr[0]
	fmt.Println(strSin)
	q := []int{0}
	q = q[1:len(q)]

	channel := make(chan int)
	flagChan := make(chan bool)
	flagChan <- true
	go byteDanceFaceTry.Producer(channel)
	go byteDanceFaceTry.Customer(channel, "cat:", flagChan)
	go byteDanceFaceTry.Customer(channel, "dog:", flagChan)
	for {

	}
	//byteDanceFaceTry.Customer(channel, flagChan)

	/*
		inputSlice := []int{1, 2, 3}
		ans := firstHomeWork.PlusOne(inputSlice)
		fmt.Println(ans)
		//int转string
		strconv.Itoa(1)
		inputSlice = []int{3, 1, 2}
		byteDanceFaceTry.Quick_sort(inputSlice, 0, len(inputSlice))
		fmt.Println(inputSlice)
		inputSlice = []int{3, 1, 2}
		byteDanceFaceTry.Merge_sort(inputSlice, make([]int, len(inputSlice)), 0, len(inputSlice))
		fmt.Println(inputSlice)
		fmt.Println(strconv.Atoi("1"))
		fmt.Println(twoSum(inputSlice, 5))
	*/
	/*
			["LRUCache","get","put","get","put","put","get","get"]
			[[2],[2],[2,6],[1],[1,5],[1,2],[1],[2]]

			["LRUCache","put","put","get","put","get","put","get","get","get"]
			[[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]

		fmt.Println("LRUTest")
		//cmdStr :=[]string{"LRUCache","put","put","get","put","get","put","get","get","get"}
		//param := [][]int{{2},{1,1},{2,2},{1},{3,3},{2},{4,4},{1},{3},{4}}
		cmdStr := []string{"LRUCache", "get", "put", "get", "put", "put", "get", "get"}
		param := [][]int{{2}, {2}, {2, 6}, {1}, {1, 5}, {1, 2}, {1}, {2}}
		var cache week2.LRUCache
		for i, cm := range cmdStr {
			switch cm {
			case "LRUCache":
				cache = week2.Constructor(param[i][0])
				fmt.Println("null")
				break
			case "get":
				fmt.Println(cache.Get(param[i][0]))
				break
			case "put":
				cache.Put(param[i][0], param[i][1])
				fmt.Println("null")
				break
			}
		}
	*/
}

func twoSum(nums []int, target int) []int {
	pair := [][]int{} //go没有tuple用数组代替
	for i, val := range nums {
		pair = append(pair, []int{val, i})
	}
	sort.Slice(pair, func(i, j int) bool {
		return pair[i][0] < pair[j][0]
	})
	l, r := 0, len(nums)-1
	for l < r {
		sumVal := pair[l][0] + pair[r][0]
		if sumVal == target {
			return []int{Global.Min(pair[l][1], pair[r][1]), Global.Max(pair[l][1], pair[r][1])}
		} else if sumVal > target {
			r--
		} else {
			l++
		}
	}
	return nil
}

func maxSlidingWindow(nums []int, k int) []int {
	//懒惰删除
	var ans []int
	q := new(myHeap)
	for i := 0; i < k-1; i++ {
		heap.Push(q, []int{nums[i], i})
	}
	for i := k - 1; i < len(nums); i++ {
		heap.Push(q, []int{nums[i], i})
		for (*q)[0][1] <= i-k {
			heap.Pop(q)
		}
		ans = append(ans, (*q)[0][0])
	}
	return ans
}
