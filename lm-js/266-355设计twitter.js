/*
355. 设计推特
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

*/

/*
    思路：设计两个列表数组。
        1.推文列表数组TweetList，创建推文时，从头部插入数组。即时间顺序
        2.关注列表数组followLists。以关注者为下标，被关注者放入到数组中。每次都可以知道当前用户关注了哪些人。
        设计方法：
        postTweet：头部插入即可
        getTweet:
            遍历推文
            1.该用户发布的TweetList[userId]==TweetLists[i][0]
            2.该用户关注的人发布的
            followLists[userId].indexOf(TweetLists[i][0]) 
            踩坑，注意返回数组长度为10时即可返回。
        follow：
        如果当前用户followerId没有关注别人，则followLists[followerId]=[followeeId]
        有关注的人，则followLists[followerId].push(followeeId)

        unfollow：
        遍历当前用户的关注人员
        followLists[followerId]
        如果关注人员中有要取消关注的人员followeeId
        则数组followLists[followerId]删除splice(i,1)当前人员

        踩坑：注意是否为空，初始化为空数组等细节控制

    题解：别人的优化和思考有点厉害。
        1.关注和取消关注。只需要控制某个人关注列表。为了删除和添加方便。可以用哈希表
        2.每个人的推文和id也可以放在哈希表中。
        对于每个推文的操作，只有顺序添加，因此用链表(不考虑扩容)和数组(动态，需考虑扩容)都可。
        查找最近10条推文。先把关注的人的列表拿出来。再合并，再排序（到这里我就有点迷糊了），经典的多路归并问题。可使用优先队列(超纲超纲)。基础知识不够，后面是真不太行了。


*/

/**
 * Initialize your data structure here.
 */
 var Twitter = function() {
    this.TweetLists=[]
    this.followLists=[]
};

/**
 * Compose a new tweet. 
 * @param {number} userId 
 * @param {number} tweetId
 * @return {void}
 */
Twitter.prototype.postTweet = function(userId, tweetId) {
    this.TweetLists.unshift([userId,tweetId])
};

/**
 * Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. 
 * @param {number} userId
 * @return {number[]}
 */
Twitter.prototype.getNewsFeed = function(userId) {
    let target=[]
    for(let i=0;i<this.TweetLists.length;i++){
        if(this.TweetLists[i][0]==userId||(this.followLists[userId]&&this.followLists[userId].indexOf(this.TweetLists[i][0])!=-1)){
            target.push(this.TweetLists[i][1])
        }
    }
    return target.slice(0,10)
};

/**
 * Follower follows a followee. If the operation is invalid, it should be a no-op. 
 * @param {number} followerId 
 * @param {number} followeeId
 * @return {void}
 */
Twitter.prototype.follow = function(followerId, followeeId) {
    this.followLists[followerId]?this.followLists[followerId].push(followeeId):(this.followLists[followerId]=[followeeId])
};

/**
 * Follower unfollows a followee. If the operation is invalid, it should be a no-op. 
 * @param {number} followerId 
 * @param {number} followeeId
 * @return {void}
 */
Twitter.prototype.unfollow = function(followerId, followeeId) {
    if(this.followLists[followerId])
    {
        for(let i=0;i<this.followLists[followerId].length;i++){
            if(this.followLists[followerId][i]==followeeId){
                this.followLists[followerId].splice(i,1)
            }
        }
    }
};

/**
 * Your Twitter object will be instantiated and called as such:
 * var obj = new Twitter()
 * obj.postTweet(userId,tweetId)
 * var param_2 = obj.getNewsFeed(userId)
 * obj.follow(followerId,followeeId)
 * obj.unfollow(followerId,followeeId)
 */


/*
    思考：看到了优的题解，但没抠出来优的代码。
         为什么使用链表存储推文(值得思考)
         https://leetcode-cn.com/problems/design-twitter/solution/ha-xi-biao-lian-biao-you-xian-dui-lie-java-by-liwe/
*/