给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。
具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
示例 1：
输入："abc"
输出：3
解释：三个回文子串: "a", "b", "c"
示例 2：
输入："aaa"
输出：6
解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
 
提示：
输入的字符串长度不会超过 1000 。

1. Clearfication:

2. Coding:

3. 看题解：
思路：
朴素方法：枚举出所有的回文子串，然后判断是否是回文
枚举出每一个可能的回文中心，然后用两个指针分别向左右两遍扩展，当两个指针指向的元素相同的时候就拓展，否则停止拓展
字符串为n，O(n * n) 的时间枚举出所有的子串，然后再用 O(ri - li + 1)的时间检测当前的子串是否是回文，整个算法的时间复杂度是O(n * n * n)
枚举回文中心是O(n),对于每个回文中心拓展的次数也是O(n),所以时间复杂度是O(n*n).
两层循环，考察所有子串，判断是否是回文串。时间复杂度是O(n * n * n),空间复杂度是O(1)
func countSubstrings(s string)int {
    count := 0
    
    for i := 0;i < len(s);i++ {
        for j := i;j < len(s);j++ {
            if isPalindrome(s[i:j+1]) {
                count++
            }   
        }
    }
    
    return count
}

func isPalindrome(s string)bool {
    i ,j := 0,len(s) -1 
    
    for i < j {
        if s[i] != s[j] {
            return false
        }
        i++
        j--
    }
    
    return true
}

这个for循环还是自己之前没有想到的
for i := 0;i < len(s);i++ {
    for j := i;j < len(s);j++ {
    
    }
}

初始化的时候 j := i
中间点可能落在字符身上，也可能落在两个字符之间
func countSubstrings(s string) int {
    ret := 0
    
    for i := 0;i < len(s);i++ {
        ret += check(i, i, s)
        ret += check(i, i + 1,s)
    }
    
    return ret
}

func check(i int,j int,s string) int {
    cnt := 0
    
    for i >= 0 && j < len(s) && s[i] == s[j] {
        cnt += 1
        i -= 1
        j += 1
    }
    
    return cnt
}

动态规划：
得动手画画图，字符串的问题，一般需要将字符串横着画一行，竖着画一行，然后将它转化为二维数组
1. 大问题是什么？
2. 规模小一点的子问题是什么？
3. 它们之间有什么联系？

大问题是一个子串是否是回文串，然后统计有多少个回文子串，那规模小一点的子问题呢？
一个子串是回文串，刨去相同的首尾字符，剩下的子串也必须是回文串。
i <= j，所以只用填半个table
s[i:j] 什么时候是回文？（dp[i][j] 为 true)

1. 由当个字符组成
2. 由2个字符组成，且字符相同
3. 由超过2个字符组成，首尾字符相同，且剩余子串是一个回文串
func countSubstrings(s string) int {
    count := 0
    dp := make([][]bool, len(s))
    
    for i := 0;i < len(dp);i++ {
        dp[i] = make([]bool, len(s))
    }
    
    for j := 0;j < len(s);j++ {
        for i := 0;i <= j;i++ {
            if i == j {
                dp[i][j] = true
                count++
            }else if j - i == 1 && s[i] == s[j] {
                dp[i][j] = true
                count++
            }else if j - i > 1 && s[i] == s[j] && dp[i + 1][j - 1] {
                dp[i][j] = true
                count++
            }
        }
    }
    
    return count
}

4. 复杂度分析：
朴素法：时间复杂度：O(n*n*n),空间复杂度O(1)
模拟中心法：时间复杂度：O(n*n),空间复杂度O(1)
动态规划：时间复杂度：O(n*n),空间复杂度O(n*n)

5. 总结：
5.1: 自己对这种字符串的问题有一丢丢的害怕，哈哈哈，感觉对字符串不是很熟悉

5.2: 自己想了一会没有思路，其实还是太跳跃了，最常规的解法要有的哈，输出所有的字符串，然后去判断是否是回文

5.3: 想不清楚的时候就动手画画
