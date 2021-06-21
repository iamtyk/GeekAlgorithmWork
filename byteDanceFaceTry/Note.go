package byteDanceFaceTry

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
