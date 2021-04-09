<?php
/*
 * @Descripttion: 周五啦啦啦
 * @Author: tacks321@qq.com
 * @Date: 2021-04-09 14:26:47
 * @LastEditTime: 2021-04-09 15:39:53
 */



/*
 * @lc app=leetcode.cn id=392 lang=php
 *
 * [392] 判断子序列
 *
 * https://leetcode-cn.com/problems/is-subsequence/description/
 *
 * algorithms
 * Easy (50.75%)
 * Likes:    427
 * Dislikes: 0
 * Total Accepted:    114.9K
 * Total Submissions: 225.7K
 * Testcase Example:  '"abc"\n"ahbgdc"'
 *
 * 给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
 * 
 * 
 * 字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。
 * （例如，"ace"是"abcde"的一个子序列，而"aec"不是）。
 * 
 * 进阶：
 * 
 * 如果有大量输入的 S，称作 S1, S2, ... , Sk 其中 k >= 10亿，你需要依次检查它们是否为 T
 * 的子序列。在这种情况下，你会怎样改变代码？
 * 
 * 致谢：
 * 
 * 特别感谢 @pbrother 添加此问题并且创建所有测试用例。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "abc", t = "ahbgdc"
 * 输出：true
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s = "axc", t = "ahbgdc"
 * 输出：false
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 0 
 * 0 
 * 两个字符串都只由小写字符组成。
 * 
 * 
 */

// @lc code=start
class Solution1 {

    /**
     * @param String $s
     * @param String $t
     * @return Boolean
     */
    // 双指针
    // 执行用时：8 ms, 在所有 PHP 提交中击败了76.00%的用户
    // 内存消耗：15.2 MB, 在所有 PHP 提交中击败了68.00%的用户
    function isSubsequence($s, $t) {
        $lenS = strlen($s);
        $lenT = strlen($t);
        // 双指针
        $iS   = 0;
        $iT   = 0;
        // 从前向后，依次进行一次比较
        while($iS < $lenS && $iT < $lenT) {
            // 相等就同时向后移动
            if($s[$iS] == $t[$iT]) {
                $iS++;
                $iT++;
            } else {
                $iT++;
            }
        }
        return $iS == $lenS; // 最后查看是否移动到末尾

    }
}


class Solution2 {

    /**
     * @param String $s
     * @param String $t
     * @return Boolean
     */
    // 动态规划
    // 执行用时：44 ms, 在所有 PHP 提交中击败了12.00%的用户
    // 内存消耗：28.6 MB, 在所有 PHP 提交中击败了8.00%的用户
    function isSubsequence($s, $t) {
        $lenS = strlen($s);
        $lenT = strlen($t);

        // dp数组dp[i][j] 表示字符串t以i位置开始第一次出现字符j的位置
        $dp = [];

        // 初始化边界条件，边界填充无穷大 ，因为$dp[$lenT]，表示t并不存在字符
        for ($i=0; $i < 26; $i++) { 
            $dp[$lenT][$i] = $lenT;
        }

        // 倒序构建
        for ($i=$lenT - 1; $i >=0 ; $i--) { 
            for ($j=0; $j < 26; $j++) { 
                // 如果当前字符整好出现 就是本身位置
                if($t[$i] == chr($j + ord('a'))  ) {
                    $dp[$i][$j] = $i;
                } else {
                    // 否则使用下一行数据
                    $dp[$i][$j] = $dp[$i+1][$j];
                }
            }
        }
        
        $add = 0;
        // 遍历子序列s
        for ($i=0; $i <$lenS ; $i++) { 
            //  dp 中没有s[i]
            if($dp[$add][ord($s[$i]) - ord('a')] == $lenT) {
                return false;
            }
            // 直接跳到t中s[i]第一次出现的位置之后一位
            $add = $dp[$add][ord($s[$i]) - ord('a')] + 1;
        }

        
        return true;


    }
}

$obj = new Solution2();
var_dump($obj->isSubsequence('axc','ahbgdc'));

 

// @lc code=end

/*
【题目】
    字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。
    （例如，"ace"是"abcde"的一个子序列，而"aec"不是）

    这个题看起来也是挺经典吧，就是把一个源字符串，里面随便删除一些字符，剩下的字符重新成为子序列，

    看起啦好难呜呜呜~动态规划
【解析】
    先来看一下双指针，双指针的思路还是挺好理解的。

    【双指针】
        s 子序列
        t 元序列
        两个指针，分别指向序列的开头，然后进行依次顺位比较。
            当两个指针对应的值相等（每次贪心的比较第一个相等的t的位置元素）
                两个指针分别向后移动
            当s的指针的值不能与t的指针
                那么移动t的指针向后
        如果最后s的指针可以移动到最后
            那么就是OK

        时间复杂度：O(n+m)
        空间复杂度：O(1)
        

    来看看动态规划到底有多难
    【动态规划】
        前面的双指针的做法，可以看到有大量的时间在用于找到t中下一个匹配的字符

        首先写出来动态规划方程
        动态规划DP数组，就是对t进行预处理！
        然后可以对其进行跳跃处理
        
        这样如果有大量的s这样字符串需要判断，则可以大大减少这种时间。

    状态转移
        创建一个二维数组，每一行i是字符的位置，每一列是a~z（两个字符串都只由小写字符组成）
        f[i][j] 表示字符串t中，从位置字符i开始往后字符j第一次出现的位置。
        如果t中i位置就是字符j，那么f[i][j] = i    
        如果j出现在i+1往后，那么f[i][j] = f[i+1][j] (也就是沿用下一行的数据)

        所以倒着进行动态规划DP数组的构建

        通过这样的构建，我们可以利用f数组，每次O(1)时间复杂度跳转到下一个位置。

    虽然看完还是有点蒙蔽，但是也算是真正看明白动态规划题的解题思路。
        

*/