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
输出: fals

1. Clarfication:
题目中说：
给定一个非空字符串s和一个包含非空单词的列表 wordDict,判定s是否可以被空格拆分为一个或多个在字典中出现的单词。
第三个 case 没有看懂：
s = "catsanddog",wordDict=["cats","dog","sand","and","cat"]
后来仔细想了下：
cats  and dog 用空格只能分隔成三个，分隔不出来多的了？
题目里面的说明又说：
拆分时可以重复使用字典中的单词 ？ 有点懵了？ 可以重复使用单词，为什么第三个case 不行？
那如果不考虑第三个case，前面两个你会怎么思考呢？
也没有太多的思路

2. 看题解：
看完了题解，明白了为什么第三个case 不行了
s = "catsandog", 可以被空格分为 cats and og，og 是不合法的
自己没有看清楚，以为自己的思路是对的，其实是自己没有看清楚
动态规划5步曲：
1. 动态方程定义：
       dp[i] 表示字符串s前i个字符组成的字符串 s[0...i-1] 是否能被空格拆分成若干个字典中出现的单词
2. 动态方程
     dp[i]  = dp[j] && check(s[j..i-1])
3. 动态方程初始化 
    dp[0] = true

func wordBreak(s string,wordDict []string)bool {
    wordDictSet := make(map[string]bool)
    for _,w := range wordDict {
        wordDictSet[w] = true
    }
    dp := make([]bool, len(s) + 1)
    dp[0] = true
    for i := 1;i <= len(s);i++ {
        for j := 0;j < i;j++ {
            if dp[j] && wordDictSet[s[j:i]] {
                dp[i] = true
                break
            }
        }
    }
    
    return dp[len(s)]
}

3. 复杂度分析：
时间复杂度：O(n*n)
空间复杂度：O(n)

4. 总结：
4.1: 理解题意还是蛮重要的，你以为的不一定是对的，要去分析

4.2: 里面的 for j := 0;j < i;j++ ,j < i这个还是值得学习和思考的，这个地方也是我不会的地方
