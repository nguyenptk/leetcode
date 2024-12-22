// https://leetcode.com/problems/design-twitter/
package medium

import "container/heap"

type Tweet struct {
	count      int
	tweetId    int
	followeeId int
	index      int
}

type TweetHeap []*Tweet

func (h TweetHeap) Len() int            { return len(h) }
func (h TweetHeap) Less(i, j int) bool  { return h[i].count < h[j].count }
func (h TweetHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *TweetHeap) Push(x interface{}) { *h = append(*h, x.(*Tweet)) }
func (h *TweetHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Twitter struct {
	count     int
	tweetMap  map[int][]*Tweet // userId: list of [count, tweetId]
	followMap map[int]map[int]bool
}

func ConstructorTwitter() Twitter {
	return Twitter{
		tweetMap:  make(map[int][]*Tweet),
		followMap: make(map[int]map[int]bool),
	}
}

func (f *Twitter) PostTweet(userId int, tweetId int) {
	if _, ok := f.tweetMap[userId]; !ok {
		f.tweetMap[userId] = make([]*Tweet, 0)
	}
	f.tweetMap[userId] = append(f.tweetMap[userId], &Tweet{
		count:   f.count,
		tweetId: tweetId,
	})
	f.count -= 1
}

func (f *Twitter) GetNewsFeed(userId int) []int {
	res := make([]int, 0)
	h := TweetHeap{}

	if _, ok := f.followMap[userId]; !ok {
		f.followMap[userId] = make(map[int]bool)
	}
	f.followMap[userId][userId] = true
	for followeeId := range f.followMap[userId] {
		if _, ok := f.tweetMap[followeeId]; ok {
			index := len(f.tweetMap[followeeId]) - 1
			tweet := f.tweetMap[followeeId][index]
			heap.Push(&h, &Tweet{
				count:      tweet.count,
				tweetId:    tweet.tweetId,
				followeeId: followeeId,
				index:      index - 1})
		}
	}

	for len(h) > 0 && len(res) < 10 {
		tweet := heap.Pop(&h).(*Tweet)
		res = append(res, tweet.tweetId)
		if tweet.index >= 0 {
			nextTweet := f.tweetMap[tweet.followeeId][tweet.index]
			heap.Push(&h, &Tweet{count: nextTweet.count,
				tweetId:    nextTweet.tweetId,
				followeeId: tweet.followeeId,
				index:      tweet.index - 1})
		}
	}
	return res
}

func (f *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := f.followMap[followerId]; !ok {
		f.followMap[followerId] = make(map[int]bool)
	}
	f.followMap[followerId][followeeId] = true
}

func (f *Twitter) Unfollow(followerId int, followeeId int) {
	if _, ok := f.followMap[followerId]; ok {
		delete(f.followMap[followerId], followeeId)
	}
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
