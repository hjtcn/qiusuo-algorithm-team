/*
139. 单词拆分
给定一个非空字符串 s 和一个包含非空单词的列表 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。

说明：

拆分时可以重复使用字典中的单词。
你可以假设字典中没有重复的单词。
示例 1：

输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
示例 2：

输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
     注意你可以重复使用字典中的单词。
示例 3：

输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false

*/

/*
    思路：架子我搭好了，但是我推不出来状态转移方程。好吧，还是0.
    思路：1.dp数组的意义
          dp[sLen]长度为[0,sLen]的字符串是否拆分形成的单词在字典中都出现
         2.状态转移方程
         如"leetcode"的一部分  字典中含有单词    
                leet    ==     leet

                code    ==     code
        推导i和j的关系
         如果(s.substr(i-word[j].length,i)==word[j])&&dp[i-word[j].length]
         理解一下：如果s切割的字符串==字典中的某个单词，且s从(0~i-word[j].length)之前的字符串拆分，也会在字典中出现。
         则dp[i]=true

         3.dp初始化
         dp[0]=1

         4.遍历顺序
          对于背包来说
         求组合数是外层遍历物品，内层遍历背包
         求排列数是外层遍历背包，内层遍历物品
         该题求是否都出现过，所以无所谓排列数还是组合数，都可以的。

         5.举例。
*/

var wordBreak = function (s, wordDict) {
    let len = wordDict.length, lenS = s.length
    let dp = Array(lenS + 1).fill(0)
    dp[0] = 1
    for (let i = 1; i <= lenS; i++) {
        for (let j = 0; j < len; j++) {
            if (i - wordDict[j].length >= 0) {
                // console.log(s.slice(i - wordDict[j].length, i), j)
                if (s.slice(i - wordDict[j].length, i) == wordDict[j] && dp[i - wordDict[j].length]) {
                    //    console.log(i,j)
                    dp[i] = true
                }
            }
        }
    }
    return dp[lenS]

};

/*
    时间复杂度：O(N*M)
    空间复杂度：O(N)
*/

/*
    思考：对于字符串的拆分，不太擅长，尤其是关于索引切割的时候。
*/