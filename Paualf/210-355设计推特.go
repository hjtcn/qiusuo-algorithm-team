设计一个简化版的推特(Twitter)，可以让用户实现发送推文，关注/取消关注其他用户，能够看见关注人（包括自己）的最近十条推文。你的设计需要支持以下的几个功能：

postTweet(userId, tweetId): 创建一条新的推文
getNewsFeed(userId): 检索最近的十条推文。每个推文都必须是由此用户关注的人或者是用户自己发出的。推文必须按照时间顺序由最近的开始排序。
follow(followerId, followeeId): 关注一个用户
unfollow(followerId, followeeId): 取消关注一个用户
示例:

Twitter twitter = new Twitter();

// 用户1发送了一条新推文 (用户id = 1, 推文id = 5).
twitter.postTweet(1, 5);

// 用户1的获取推文应当返回一个列表，其中包含一个id为5的推文.
twitter.getNewsFeed(1);

// 用户1关注了用户2.
twitter.follow(1, 2);

// 用户2发送了一个新推文 (推文id = 6).
twitter.postTweet(2, 6);

// 用户1的获取推文应当返回一个列表，其中包含两个推文，id分别为 -> [6, 5].
// 推文id6应当在推文id5之前，因为它是在5之后发送的.
twitter.getNewsFeed(1);

// 用户1取消关注了用户2.
twitter.unfollow(1, 2);

// 用户1的获取推文应当返回一个列表，其中包含一个id为5的推文.
// 因为用户1已经不再关注用户2.
twitter.getNewsFeed(1);

1. Clarification:
关注与取消关注

获取关注者的 list 列表信息

Q:关系怎么存储呢
A：列表想到了用list 进行存储，是否follower 觉得可以使用map 进行存储，map[id]id

没啥思路，数据流画不出来。。。


2. 看题解：

说实话看题解里面说到获取最近10条的时候用到什么大顶堆什么的，好像有点麻烦呦
看了下面的题解，在全局列表里面倒叙在自己关注列表里面倒叙获取10条

回顾下，这道题自己为啥没写出来，
50%是害怕，30%是懒，20%可能还是早上有点饿了 。。。

复盘下，关注关系自己是想出来了，列表的获取自己没想出来。。。

自己想的是用数组存储性能不好？最好的方式是用链表
但是如果你连基本的数据流都实现不了，关注性能又有什么用呢？

是不是想太多了，想把数据流跑通，然后再去优化是不是会更好一点
type Twitter struct {
	Tweets []int
	UserTweets map[int][]int
	Follows map[int][]int
	IsFollowMy map[int]bool
}


/** Initialize your data structure here. */
func Constructor() Twitter {
	// 每一次实例化的时候，都重新分配一次，这样不会造成示例重复

	var Tweets  []int

	// 某用户发的某条推特
	var UserTweets = make(map[int][]int)

	// 某用户关注了哪些用户
	var Follows = make(map[int][]int)

	var IsFollowMy = make(map[int]bool)

	t := Twitter{
		Tweets:Tweets,
		UserTweets:UserTweets,
		Follows: Follows,
		IsFollowMy: IsFollowMy,
	}
	return t
}


/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int)  {
	// 每个人每次发推特，都记录到一个地方
	this.Tweets = append(this.Tweets,tweetId)
	// 某个用户发了推特，存到自己推特列表里
	this.UserTweets[userId] = append(this.UserTweets[userId],tweetId)
}


/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	fs := this.Follows[userId] // 先获取该用户的关注列表
	var allTweets []int
	for _,v := range fs {
		// 把关注列表的人的所有推特都集中起来
		allTweets = append(allTweets,this.UserTweets[v]...)
	}
	if !this.IsFollowMy[userId] {
		// 如果自己没有关注自己，那么也需要把自己发的推特加到一起
		allTweets = append(allTweets,this.UserTweets[userId]...)
	}
	var sortTweets []int
	aTLen := len(this.Tweets)
	s := 0
	// 按照发的推特顺序进行倒序排序
	for i:=aTLen-1;i>=0;i-- {
		if s >= 10 {
			break
		}
		for _,n := range allTweets {
			
			// 只取 10条数据
			if this.Tweets[i] == n && s < 10{
				s++
				sortTweets = append(sortTweets,n)
			}
		}
	}

	return sortTweets
}


/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int)  {
	// 如果自己关注了自己，标记一下
	if followerId == followeeId {
		this.IsFollowMy[followerId] = true
	}
	
	// 下面是判断这人是否关注了，如果已经关注了，那么就不再关注了
	var isFed bool
	for _,v := range this.Follows[followerId] {
		if v == followeeId {
			isFed = true
		}
	}
	if !isFed {
		this.Follows[followerId] = append(this.Follows[followerId],followeeId)
	}
}


/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int)  {
	// 如果自己取关了自己，标记一下
	if followeeId == followerId {
		this.IsFollowMy[followerId] = false
	}
	
	// 去掉自己关注列表里那个被关注的人
	var temp []int
	for _,v := range this.Follows[followerId] {
		if v != followeeId {
			temp = append(temp,v)
		}
	}
	this.Follows[followerId] = temp
}
/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */

3.复杂度：
时间复杂度：O(nxn)
空间复杂度：O(n)

4.总结：
4.1: 不要过早关注性能，一开始优先关注实现与数据流的流转
