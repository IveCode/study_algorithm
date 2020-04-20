package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {

	//["Twitter","postTweet","getNewsFeed","follow","postTweet","getNewsFeed","unfollow","getNewsFeed"]
	//[[],[1,5],[1],[1,2],[2,6],[1],[1,2],[1]]
	//obj := Constructor()
	//obj.PostTweet(1, 5)
	//fmt.Println(obj.GetNewsFeed(1))
	//obj.Follow(1, 2)
	//obj.PostTweet(2, 6)
	//fmt.Println(obj.GetNewsFeed(1))
	//obj.Unfollow(1, 2)
	//fmt.Println(obj.GetNewsFeed(1))

	//["Twitter","postTweet","getNewsFeed","follow","getNewsFeed","unfollow","getNewsFeed"]
	//[[],[1,1],[1],[2,1],[2],[2,1],[2]]
	//obj := Constructor()
	//obj.PostTweet(1, 1)
	//fmt.Println(obj.GetNewsFeed(1))
	//obj.Follow(2, 1)
	//fmt.Println(obj.GetNewsFeed(2))
	//obj.Unfollow(2, 1)
	//fmt.Println(obj.GetNewsFeed(2))

	//["Twitter","postTweet","follow","getNewsFeed"]
	//[[],[1,5],[1,1],[1]]
	//obj := Constructor()
	//obj.PostTweet(1, 5)
	//obj.Follow(1, 1)
	//fmt.Println(obj.GetNewsFeed(1))

	//["Twitter","postTweet","getNewsFeed","follow","postTweet","getNewsFeed","unfollow","getNewsFeed"]
	//[[],[1,5],[1],[1,2],[2,6],[1],[1,2],[1]]
	obj := Constructor()
	obj.PostTweet(1, 5)
	fmt.Println(obj.GetNewsFeed(1))
	obj.Follow(1, 2)
	obj.PostTweet(2, 6)
	fmt.Println(obj.GetNewsFeed(1))
	obj.Unfollow(1, 2)
	fmt.Println(obj.GetNewsFeed(1))
}

type Tweet struct {
	TweetId   int
	Timestamp int64
}

type Tweets []Tweet

func (p Tweets) Len() int           { return len(p) }
func (p Tweets) Less(i, j int) bool { return p[i].Timestamp > p[j].Timestamp }
func (p Tweets) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type UserInfo struct {
	UserId       int
	Tweets       Tweets
	FollowUserId map[int]struct{}
}

func NewUserInfo(userId int) *UserInfo {
	return &UserInfo{UserId: userId, Tweets: Tweets{}, FollowUserId: map[int]struct{}{}}
}

type Twitter struct {
	Value map[int]*UserInfo
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{Value: map[int]*UserInfo{}}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	userInfo, ok := this.Value[userId]
	if !ok {
		userInfo = NewUserInfo(userId)
		this.Value[userId] = userInfo
	}

	userInfo.Tweets = append(userInfo.Tweets, Tweet{
		TweetId:   tweetId,
		Timestamp: time.Now().UnixNano(),
	})
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {

	userInfo, ok := this.Value[userId]
	if !ok {
		return []int{}
	}
	var tweets Tweets
	tweets = append(tweets, userInfo.Tweets...)
	for userId, _ := range userInfo.FollowUserId {
		oTweets := this.getTweets(userId)
		if len(oTweets) > 0 {
			tweets = append(tweets, oTweets...)
		}
	}
	sort.Sort(tweets)
	var tweetIds []int
	for i, tweet := range tweets {
		if i >= 10 {
			break
		}
		tweetIds = append(tweetIds, tweet.TweetId)
	}

	return tweetIds
}

func (this *Twitter) getTweets(userId int) Tweets {
	userInfo, ok := this.Value[userId]
	if ok {
		return userInfo.Tweets
	}
	return Tweets{}
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	if followerId == followeeId {
		return
	}
	userInfo, ok := this.Value[followerId]
	if !ok {
		userInfo = NewUserInfo(followerId)
		this.Value[followerId] = userInfo
	}
	userInfo.FollowUserId[followeeId] = struct{}{}
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	userInfo, ok := this.Value[followerId]
	if !ok {
		return
	}
	delete(userInfo.FollowUserId, followeeId)
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
