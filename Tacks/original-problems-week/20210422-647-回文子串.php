<?php
/*
 * @Descripttion: 周四啦，加油干
 * @Author: tacks321@qq.com
 * @Date: 2021-04-22 09:48:52
 * @LastEditTime: 2021-04-22 20:06:43
 */


/*
 * @lc app=leetcode.cn id=647 lang=php
 *
 * [647] 回文子串
 *
 * https://leetcode-cn.com/problems/palindromic-substrings/description/
 *
 * algorithms
 * Medium (64.71%)
 * Likes:    522
 * Dislikes: 0
 * Total Accepted:    86.1K
 * Total Submissions: 132.3K
 * Testcase Example:  '"abc"'
 *
 * 给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。
 * 
 * 具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入："abc"
 * 输出：3
 * 解释：三个回文子串: "a", "b", "c"
 * 
 * 
 * 示例 2：
 * 
 * 输入："aaa"
 * 输出：6
 * 解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 输入的字符串长度不会超过 1000 。
 * 
 * 
 */

// @lc code=start
class Solution1 {

    /**
     * @param String $s
     * @return Integer
     */
    // 中心扩展法
    // 执行用时：32 ms, 在所有 PHP 提交中击败了100.00%的用户
    // 内存消耗：15.2 MB, 在所有 PHP 提交中击败了100.00%的用户
    function countSubstrings($s) {
        $len = strlen($s);

        $res = 0;
        for ($i=0; $i < $len; $i++) { 
            $res += $this->_helper($s, $len, $i, $i);  // 奇数
            $res += $this->_helper($s, $len, $i, $i+1);// 偶数
        }
        return $res;
    }

    private function _helper($s, $len, $left, $right) {
        $count = 0;
        while($left >=0 && $right < $len && $s[$left] == $s[$right]) {
            $left--;
            $right++;
            $count++;
        }
        return $count;
    }
}

class Solution2 {

    /**
     * @param String $s
     * @return Integer
     */
    // 动态规划
    // 执行用时：180 ms, 在所有 PHP 提交中击败了18.75%的用户
    // 内存消耗：35 MB, 在所有 PHP 提交中击败了37.50%的用户
    function countSubstrings($s) {
        $len = strlen($s);
        $count = 0; // 回文串的数量
        for ($i=$len-1; $i>=0 ; $i--) { 
            for ($j=$i; $j < $len; $j++) { 
                // 如果构不成回文串就跳过
                if($s[$i] != $s[$j]) {
                    continue;
                }
                
                $dp[$i][$j] = (($j-$i) <=2 )|| $dp[$i+1][$j-1];
                // 如果可以构成回文串
                if($dp[$i][$j]) {
                    $count++;
                }
            }
        }
        return $count;
    }

    // 优化
    
}
// @lc code=end

/*
【题意】
    回文子串，就是从已知的字符串中，找到所有的子串，判断是否时回文，同时不同位置的字符串，也是不同的回文子串
【解析】
    暴力法
        枚举每一个子串     O(n*n)
        子函数判断是否回文 O(n)

        这个只是一个思路，基本上都会超时，因为时间复杂度已经上升到 O(n*n*n)

    1、中心扩散法
        ①先找到一个中心（寻找中心）
        ②然后往外扩展（扩展函数）

        但是考虑回文串的偶数的情况？
            =>单数中心，双数中心

        伪代码

            ①寻找中心
                单数中心
                    for i
                        扩展函数(i,i)
                双数中心
                    for i
                        扩展函数(i,i+1)
            ②扩展函数
                left right两个指针
                边界判断
                    count 计数
                    扩散条件 s[l] == s[r]
                        l--
                        r++
                        count++
                返回 count

    2、动态规划
        定义dp[][]二维数组， 

        字符串s[i...j] 是否是回文串，如果是那么dp[i][j]=true 如果不是dp[i][j]=false
        
        如果我们知道dp[i+1][j-1]那么，我们如何判断dp[i][j]
            如果s[i]=s[j]，那么就说明只要dp[i+1][j-1]是回文串，那么dp[i][j]就是回文串（有点往里面缩小的感觉）
            如果s[i]!=s[j]，那么，dp[i][j] 就肯定不是
        但是还有一种情况 （感觉题解区都是大神，怎么就能想这么细节，比如cec ）
            也就是i和j非常近的情况，j-i<=2，我们也可以理解为回文
                比如cec 当i=0 j=2 的时候 j-i=2 是回文
                比如cc  当i=0 j=1 的时候 j-i=1 是回文
                比如c   当i=0 j=0 的时候 j-i=0 是回文

        动态转移方程
            dp[i][j] = j - i <= 2 || dp[i + 1][j - 1];
        
        

*/