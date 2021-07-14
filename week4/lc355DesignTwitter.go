package week4

import "container/heap"

type Tweet struct {
	id     int
	timeID int
	next   *Tweet
}

type twitterHeap []Tweet

func (h *twitterHeap) Less(i, j int) bool {
	return (*h)[i].timeID > (*h)[j].timeID
}

func (h *twitterHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *twitterHeap) Len() int {
	return len(*h)
}

func (h *twitterHeap) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *twitterHeap) Push(v interface{}) {
	*h = append(*h, v.(Tweet))
}

type Twitter struct {
	userTwitterMap map[int]*Tweet
	userFollowMap  map[int]map[int]int //用户关注列表 userID->map[followID]
	timeID         int                 //时间id
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		userTwitterMap: map[int]*Tweet{},
		userFollowMap:  map[int]map[int]int{},
	}
}

/** Compose a new tweet. */
func (t *Twitter) PostTweet(userId int, tweetId int) {
	t.timeID++
	tweetData := &Tweet{
		id:     tweetId,
		timeID: t.timeID,
	}
	oldData := t.userTwitterMap[userId]
	tweetData.next = oldData
	t.userTwitterMap[userId] = tweetData
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (t *Twitter) GetNewsFeed(userId int) []int {

	//我的关注列表
	follow, ok := t.userFollowMap[userId]
	//发送的内容中要包含自己的，把自己当成是关注的对象，加进去
	if !ok {
		follow = map[int]int{}
	}

	follow[userId] = 1
	//建堆
	myHeap := new(twitterHeap)
	for followID := range follow {
		tweetData, ok := t.userTwitterMap[followID]
		if !ok {
			continue
		}
		ptlTweet := tweetData
		//入堆
		index := 1
		for ptlTweet != nil {
			heap.Push(myHeap, *ptlTweet)
			ptlTweet = ptlTweet.next
			index++
			//不超过10条
			if index > 10 {
				break
			}
		}
	}

	//出堆
	var ans []int
	for i := 0; i < 10; i++ {
		if myHeap.Len() == 0 {
			break
		}

		heapPop := heap.Pop(myHeap).(Tweet)
		ans = append(ans, heapPop.id)
	}
	return ans
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (t *Twitter) Follow(followerId int, followeeId int) {
	follMap, ok := t.userFollowMap[followerId]
	if !ok {
		follMap = map[int]int{}
		t.userFollowMap[followerId] = follMap
	}
	follMap[followeeId] = 1
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (t *Twitter) Unfollow(followerId int, followeeId int) {
	follMap, ok := t.userFollowMap[followerId]
	if !ok {
		return
	}
	delete(follMap, followeeId)
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
