给你一个字符串 s ，字符串的「能量」定义为：只包含一种字符的最长非空子字符串的长度。
请你返回字符串的能量。
示例 1：
输入：s = "leetcode"
输出：2
解释：子字符串 "ee" 长度为 2 ，只包含字符 'e' 。
示例 2：
输入：s = "abbcccddddeeeeedcba"
输出：5
解释：子字符串 "eeeee" 长度为 5 ，只包含字符 'e' 。
示例 3：
输入：s = "triplepillooooow"
输出：5
示例 4：
输入：s = "hooraaaaaaaaaaay"
输出：11
示例 5：
输入：s = "tourist"
输出：1
 
提示：
1 <= s.length <= 500
s 只包含小写英文字母。
1. Clearfication:
  连续最长的字符串
    使用 i,j 两个指针进行比较
    ret := 0
    i:0 -> len(s) - 1
         tmp := 0
         j: i + 1,len(s) - 1
              if s[i] == s[j] {
                  tmp++
                  i++
                  j++
               }else {
                   i = j
                   break
               }
        
    return ret
    想要实现的是当不想等的时候，i = j 一直超时

2. Coding:
/*
    连续最长的字符串
    使用 i,j 两个指针进行比较
    ret := 0
    i:0 -> len(s) - 1
         tmp := 0
         j: i + 1,len(s) - 1
              if s[i] == s[j] {
                  tmp++
                  i++
                  j++
               }else {
                   i = j
                   break
               }
        
    return ret
    想要实现的是当不想等的时候，i = j 一直超时
*/
func maxPower(s string) int {
    ret := 0
    for i := 0;i <= len(s) - 1; {
        tmp := 0
        for j := i + 1;j <= len(s) - 1;j++ {
            fmt.Println(i,j)
            if s[i] == s[j] {
                tmp++
            }else {
                i = j
                break
            }
        }
    }
    return ret
}
上面代码死循环了，i 一直是 7 7 7 7 7 退出不出去了
i++ 这条语句，没有想清楚放到哪里。。。

3. 看题解：
看了题解发现不需要两次循环，遍历一次然后更新tmp 值就可以了
func maxPower(s string) int {
    n := len(s)
    ret := 1
    currentLen := 1
    
    for i := 1;i < n;i++ {
        if s[i] == s[i - 1] {
            currentLen += 1   
        }else {
            if currentLen > ret {
                ret = currentLen
            }
            currentLen = 1
        }
    }
    
    if currentLen > ret {
        ret = currentLen
    }
    
   return ret
}
func maxPower(s string) int {
    ret := 1
    left,right := 0,1
    for right < len(s) {
        if s[right] != s[left] {
            ret = max(right - left, ret)
            left = right
        }
        right += 1
    }
    
    return max(right - left, ret)
}
func max(x,y int) int {
    if x > y {
        return x
    }
    
    return y
}

4. 复杂度分析：
时间复杂度：O(n): 遍历整个字符串
空间复杂度：O(1)

5. 总结：
5.1: 以前感觉 for 循环很简单，慢慢发现用好还是不简单的
