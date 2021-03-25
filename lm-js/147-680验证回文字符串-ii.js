// 680. 验证回文字符串 Ⅱ
// 给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。

// 示例 1:

// 输入: "aba"
// 输出: True
// 示例 2:

// 输入: "abca"
// 输出: True
// 解释: 你可以删除c字符。
// 注意:

// 字符串只包含从 a-z 的小写字母。字符串的最大长度是50000。

/*
 * @lc app=leetcode.cn id=680 lang=javascript
 *
 * [680] 验证回文字符串 Ⅱ
 */
/*
    思路：双指针。若左右指针指向元素相等，则left++,right--
                不想等。则要么左边删一个，判断剩余字符串是否是回文
                        要么右边删一个，判断剩余字符串是否是回文

*/
// @lc code=start
/**
 * @param {string} s
 * @return {boolean}
 */
 var validPalindrome = function(s) {
    s=s.split('')
    let isPalindrome=(s)=>{
        return s.join('')==s.reverse().join('')
    }
    let left=0,right=s.length-1
    while(left<right){
        if(s[left]==s[right]){
            left++
            right--
        }
        else{
            if(isPalindrome(s.slice(left+1,right+1))){
                return true
            }
            else if(isPalindrome(s.slice(left,right))){
                return true
            }
            else{
                return false
            }
            //可合并为
            /**
               return isPalindrome(s.slice(left+1,right+1))||isPalindrome(s.slice(left,right))
             */
        }
    }
    return true
};

/*
    时间复杂度：O(N)
    空间负责度：O(N)

*/

//下面的方法无法解此题，因为字符串过长，构建二维数组会出现内存溢出的问题。
//时间复杂度：O(N^2),空间复杂度：O(N^2)
//个人思索：只能删除一个，那要是操作多个呢？上面的方法就用不了了。
//后来百度了一下，发现它的拓展是典型的动态规划问题。这里提前尝试一下吧。

/*dp[i][j]代表从i到j所需的最小操作数。
  已知'x'是回文串，左右两边都加一个‘a'也是一个回文串。
  因此：
  1.当s[i]==s[j],,,则dp[i][j]=dp[i+1][j-1](这个操作就仿佛左加了一个s[i],右加了一个s[j])，因此dp[i][j]的值不变
  2.当s[i]!=s[j]时，此时需要变成更小的子问题。有两种方案可操作:(1)左边加个s[j]字符,(2)右边加一个s[i]字符
  因此向更小的子问题扩展，即当前dp[i][j]=Min(在左操作，在右操作)+1.
    在左操作为：在左加一个s[j],即s[i][j-1],右指针左移一位，代表，左边加上啦。
    在右操作为：在右加一个s[i],即s[i+1][j],左指针右移一位，代表，右边加上啦
    
    总结：if(s[i]==s[j]){
            dp[i][j]=dp[i+1][j-1]
        }
        else{
            dp[i][j]=Math.min(dp[i+1][j],dp[i][j-1])+1
        }
*/
var validPalindrome = function(s) {
    let len=s.length,dp=[];
    for(let i=0;i<len;i++){
        dp[i]=new Array(len).fill(0)
    }
    for(let k=1;k<len;k++){
        for(let i=0;i+k<len;i++){
            let j=k+i
            if(s[i]==s[j]){
                dp[i][j]=dp[i+1][j-1]
            }
            else{
                dp[i][j]=Math.min(dp[i+1][j],dp[i][j-1])+1
            }
        }
    }
    if(dp[0][len-1]>1){
        return false
    }
    else{
        return true
    }
}


/*
执行出错。。。
<--- Last few GCs --->
[43:0x4699750]      465 ms: Scavenge 103.9 (129.6) -> 102.5 (145.3) MB, 13.7 / 0.0 ms  (average mu = 1.000, current mu = 1.000) allocation failure
[43:0x4699750]      482 ms: Scavenge 118.5 (145.3) -> 118.6 (145.6) MB, 6.1 / 0.0 ms  (average mu = 1.000, current mu = 1.000) allocation failure
[43:0x4699750]      490 ms: Scavenge 118.6 (145.6) -> 118.3 (148.3) MB, 7.6 / 0.0 ms  (average mu = 1.000, current mu = 1.000) allocation failure
<--- JS stacktrace --->
FATAL ERROR: Scavenger: semi-space copy Allocation failed - JavaScript heap out of memory
*/