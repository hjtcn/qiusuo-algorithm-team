/*
680. 验证回文字符串 Ⅱ
给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。

示例 1:

输入: "aba"
输出: True
示例 2:

输入: "abca"
输出: True
解释: 你可以删除c字符。

*/

/*
思路：最多删除一个字符。
    双指针。左删或者右删，字符串删除也是细节
    l=0,r=s.length-1
    if(s[l]==s[r]){
        l++
        r--
    }
    else{
        //删左为回文，返回true
        if(isPalindrome(s.slice(i+1,j+1))){
            return true
        }
        else if(isPalindrome(s.slice(i,j))){
            //删右为回文，返回true
            return true
        }
        else{
            return false
        }
    }

    
*/

var validPalindrome = function (s) {
    let isPalindrome = (str) => {
        let l = 0, r = str.length - 1
        while (l < r) {
            if (str[l] == str[r]) {
                l++
                r--
            }
            else {
                return false
            }
        }
        return true
    }
    let left = 0, right = s.length - 1
    while (left < right) {
        if (s[left] == s[right]) {
            left++
            right--
        }
        else {
            //不相等，left右移或者right左移
            if (isPalindrome(s.slice(left + 1))) {
                return true
            }
            else if (isPalindrome(s.slice(left, right - 1))) {
                return true
            }
            else {
                return false
            }
        }
    }
    return true
}

/*
    时间复杂度:O(N)
    空间复杂度:O(N)判断回文函数占据的内存
*/

/*动态规划崩溃呀，js一直超出内存。
一个灵活的动态规划题解(c++)
https://leetcode-cn.com/problems/valid-palindrome-ii/solution/c-kao-lu-shan-kge-do-more-do-better-by-wen-mu-yang/
*/