package main

import (
	"GeekAlgorithmWork/byteDanceFaceTry"
	"GeekAlgorithmWork/firstHomeWork"
	"fmt"
	"strconv"
)

func main() {
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

}
