package week2

import "container/list"

//用list包做，但是耗时特别长，应该是value.([]int)的时候出问题了
//试了一下，也没快多少
//我在leetcode断点调试时，取字典的时候就很慢，可能和list包实现的双向链表是一个环有关，头和尾是相连的
//如果自己实现头尾断开的双向链表就只要120ms，用list包就要700ms
type LRUCache struct {
	cacheList *list.List            //list包中，头前结点是尾指针，
	refectMap map[int]*list.Element //[]int{key, value}
	capactiy  int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cacheList: list.New(),
		refectMap: map[int]*list.Element{},
		capactiy:  capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	val, ok := this.refectMap[key]
	if !ok {
		return -1
	}
	this.cacheList.MoveToFront(val)
	return val.Value.([]int)[1]
}

//先插入，再删除才可以，否则键值一样的时候就直接移位了
func (this *LRUCache) Put(key int, value int) {
	if val, ok := this.refectMap[key]; ok {
		val.Value = []int{key, value}
		this.cacheList.MoveToFront(val)
		return
	}
	this.cacheList.PushFront([]int{key, value})
	this.refectMap[key] = this.cacheList.Front()
	if this.capactiy < this.cacheList.Len() {
		delete(this.refectMap, this.cacheList.Back().Value.([]int)[0])
		this.cacheList.Remove(this.cacheList.Back())
	}
}

//转long版本，可能几个转换的地方还是很多
/*
type LRUCache struct {
	cacheList *list.List//list包中，头前结点是尾指针，
	refectMap map[int64]*list.Element//[]int{key, value}
	capactiy int
}

const   constantName = 100000

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cacheList : list.New(),
		refectMap :  map[int64]*list.Element{},
		capactiy:capacity,
	}
}


func (this *LRUCache) Get(key int) int {
	val, ok := this.refectMap[int64(key)]
	if !ok{
		return -1
	}
	this.cacheList.MoveToFront(val)
	return int(val.Value.(int64)%constantName)
}


func (this *LRUCache) Put(key int, value int)  {
	if val, ok := this.refectMap[int64(key)];ok{
		val.Value=int64(key*constantName+value)
		this.cacheList.MoveToFront(val)
		return
	}
	this.cacheList.PushFront(int64( key*constantName+value))
	this.refectMap[int64(key)]=this.cacheList.Front()
	if this.capactiy < this.cacheList.Len(){
		delete(this.refectMap, this.cacheList.Back().Value.(int64)/constantName)
		this.cacheList.Remove(this.cacheList.Back())
	}
}


*/
