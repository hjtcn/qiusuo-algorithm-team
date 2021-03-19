给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
 
示例 1:
输入: s = "abcabcbb"
输出: 3 
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:
输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
示例 4:
输入: s = ""
输出: 0
提示：
• 0 <= s.length <= 5 * 104
• s 由英文字母、数字、符号和空格组成

1. Clearfication:
    找出所有子串，判断是否含有重复字符，如果没有重复字符，则计算它的长度，记录其最大长度

2. Coding:
/*
    不含有重复字符的 最长子串
    找出所有子串，判断是否含有重复字符，如果没有重复字符，则计算它的长度，记录其最大长度
*/
func lengthOfLongestSubstring(s string) int {
    if len(s) == 0 {
        return 0
    }
    
    ret := 0
    for i := 0;i < len(s);i++ {
        for j := i;j < len(s);j++ {
             if isDuplicate(s[i:j + 1]) {
                 continue
             }
             len := j - i + 1
             if len > ret {
                 ret = len
             }
        }
    }
    return ret    
}

/*
    判断字符串中是否有重复元素
*/
func isDuplicate(s string)bool {
    m := make(map[rune]int, 0)
    for _,v := range s {
        if m[v] > 0 {
            return true
        }
        m[v] = m[v] + 1
    }
    return false
}
上面解法超时了

3. 看题解：
滑动窗口，不断的向右移动：
func lengthOfLongestSubstring(s string)int {
    m := map[byte]int{}
    n := len(s)
    
    rk,ans := -1,0
    for i := 0;i < n;i++ {
        if i != 0 {
            delete(m, s[i-1])
        }
        
        for rk + 1 < n && m[s[rk + 1]] == 0 {
            m[s[rk + 1]]++
            rk++
        }
        
        ans = max(ans, rk - i + 1)
    }
    
    return ans
}

func max(x,y int) int {
    if x < y {
        return y
    }
    return x
}

func lengthOfLongestSubstring(s string) int {
    maxLen,start := 0,0
    table := [128]int{}
    
    for i,_ := range table {
        table[i] = -1
    }
    
    for i,c := range s {
        if table[c] >= start {
            start = table[c] + 1
        }
        table[c] = i
        maxLen = maxInt(maxLen,i - start + 1)
    }
    
    return maxLen
}

func maxInt(a,b int) int {
    if a > b {
        return a
    }
    return b
}
这个解法还没有看懂，这个是用到动态规划了么

4. 复杂度分析：
时间复杂度：O(N),N是字符串的长度
空间复杂度：O(N), N 是map的长度

5. 总结：
5.1: 滑动窗口还是蛮有意思的，那么对比我们找出所有的字符串去判断，滑动窗口节省了哪些时间呢？找出所有字符串的时候，当我们知道前面已经有重复字符串的时候我们还是在向后找重复子串，这样会有很多多余计算量
