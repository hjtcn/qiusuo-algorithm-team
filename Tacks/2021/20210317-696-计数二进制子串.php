<?php
/*
 * @Descripttion: 周三一过，周末还远吗
 * @Author: tacks321@qq.com
 * @Date: 2021-03-17 09:23:49
 * @LastEditTime: 2021-03-17 10:55:59
 */


/*
 * @lc app=leetcode.cn id=696 lang=php
 *
 * [696] 计数二进制子串
 *
 * https://leetcode-cn.com/problems/count-binary-substrings/description/
 *
 * algorithms
 * Easy (62.32%)
 * Likes:    338
 * Dislikes: 0
 * Total Accepted:    44.8K
 * Total Submissions: 71.7K
 * Testcase Example:  '"00110"'
 *
 * 给定一个字符串 s，计算具有相同数量 0 和 1 的非空（连续）子字符串的数量，并且这些子字符串中的所有 0 和所有 1 都是连续的。
 * 
 * 重复出现的子串要计算它们出现的次数。
 * 
 * 
 * 
 * 示例 1 :
 * 
 * 
 * 输入: "00110011"
 * 输出: 6
 * 解释: 有6个子串具有相同数量的连续1和0：“0011”，“01”，“1100”，“10”，“0011” 和 “01”。
 * 
 * 请注意，一些重复出现的子串要计算它们出现的次数。
 * 
 * 另外，“00110011”不是有效的子串，因为所有的0（和1）没有组合在一起。
 * 
 * 
 * 示例 2 :
 * 
 * 
 * 输入: "10101"
 * 输出: 4
 * 解释: 有4个子串：“10”，“01”，“10”，“01”，它们具有相同数量的连续1和0。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * s.length 在1到50,000之间。
 * s 只包含“0”或“1”字符。
 * 
 * 
 */

// @lc code=start
class Solution1 {

    /**
     * @param String $s
     * @return Integer
     */
    // 一次遍历
    // 执行用时：56 ms, 在所有 PHP 提交中击败了62.50%的用户
    // 内存消耗：15.3 MB, 在所有 PHP 提交中击败了100.00%的用户
    function countBinarySubstrings($s) {
        $result = 0; // 结果
        $pre    = 0; // 前一个字符数量
        $cur    = 1; // 当前字符数量 这里是要拿后面字符比较
        $len    = strlen($s) - 1;

        # 假如字符串是  000111
        # 如果只有 000 不符合要求， 只有1，所以不满足
        # 从头遍历一直遍历到 1 ，也就是出现不相等字符的时候 pre = 3; cur = 1; 
        # 当遍历到 0001 这个时候 cur=1,pre=3 pre>cur包含子串01  符合要求 result++
        # 当遍历到 00011 这个时候 cur=2,pre=3 pre>cur包含子串0011 符合要求 result++
        # 当遍历到 000111 这个时候 cur=3,pre=3 pre=cur包含子串000111 符合要求 result++
        for ($i=0; $i < $len; $i++) { 
            // 如果当前字符和后一个相等
            if($s[$i] == $s[$i+1]) {
                $cur++; // 当前字符数量 ++
            } else {
                // 如果当前字符和后一个不相等，那么就开始替换了
                $pre = $cur; // 前一个字符数量
                $cur = 1; // 当前字符数量
            }

            // 如果前面的字符大于后面的说明是连续子串
            if($pre >= $cur) {
                $result++;
            }

        }
        return $result;
    }
}


class Solution2 {

    /**
     * @param String $s
     * @return Integer
     */
    # 按照官方题解的做法，按照字符分组
    // 时间空间复杂度O(n)
    // 执行用时：80 ms, 在所有 PHP 提交中击败了25.00%的用户
    // 内存消耗：16.4 MB, 在所有 PHP 提交中击败了12.50%的用户
    function countBinarySubstrings($s) {
        $counts = []; // 字符分组
        $i = 0;
        $len = strlen($s);

        // 遍历字符串 统计每个字符连续出现的个数 得到counts
        while($i < $len) {
            $value = $s[$i]; // 当前值
            $count = 0; // 当前值的次数
            while($i < $len && $s[$i] == $value) {
                $i++;
                $count++;
            }
            array_push($counts, $count); // 追加到数组中
        }

        // 最终结果
        $ans = 0;
        $size= count($counts);
        for ($j=0; $j < $size; $j++) { 
            $ans += min($counts[$j], $counts[$j-1]);            
        }
        return $ans;
    }
}



class Solution3 {

    /**
     * @param String $s
     * @return Integer
     */
    // 执行用时：68 ms, 在所有 PHP 提交中击败了25.00%的用户
    // 内存消耗：15.4 MB, 在所有 PHP 提交中击败了62.50%的用户
    function countBinarySubstrings($s) {
        $i = 0;
        $len = strlen($s);
        $last= 0;// 上一个字符出现次数
        $ans = 0;// 结果

        // 遍历字符串 统计每个字符连续出现的个数 得到counts
        while($i < $len) {
            $value = $s[$i]; // 当前值
            $count = 0; // 当前值的次数
            // 显然，这里我们都进行两次遍历
            while($i < $len && $s[$i] == $value) {
                $i++;
                $count++;
            }
            $ans += min($count, $last);
            $last = $count; // 更新上一个字符数量
        }

       
        return $ans;
    }
}

// @lc code=end



/*
【题目】
    好家伙，读完之后知道啥意思，就是不知道怎么得到解法

【解析】
    看了官方题解，再来实现一下
    
    没想到看到题解，这是简单题吗？哦是挺简单的

    还是挺有意思的一个题，主要看数学思维，有点像小学奥数题

    (1) 统计相邻不同字符的出现次数 【单次遍历】
        一次遍历，采用 pre cur来记录相邻不同字符的出现次数
    (2) 字符分组法
        count来统计出来不同字符的出现次数，然后比较相邻的最小的，然后进行累加。
        时间空间复杂度O(n)
    (3) 则是对(2)的 优化


    解法(1)比较好一些
    解法(2)便于理解
*/