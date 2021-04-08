<?php
/*
 * @Descripttion: 周二的 时间过的真快 一天又一天
 * @Author: tacks321@qq.com
 * @Date: 2021-03-23 18:50:17
 * @LastEditTime: 2021-03-24 14:51:38
 */


/*
 * @lc app=leetcode.cn id=1446 lang=php
 *
 * [1446] 连续字符
 *
 * https://leetcode-cn.com/problems/consecutive-characters/description/
 *
 * algorithms
 * Easy (57.80%)
 * Likes:    15
 * Dislikes: 0
 * Total Accepted:    9.9K
 * Total Submissions: 17.1K
 * Testcase Example:  '"leetcode"'
 *
 * 给你一个字符串 s ，字符串的「能量」定义为：只包含一种字符的最长非空子字符串的长度。
 * 
 * 请你返回字符串的能量。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 输入：s = "leetcode"
 * 输出：2
 * 解释：子字符串 "ee" 长度为 2 ，只包含字符 'e' 。
 * 
 * 
 * 示例 2：
 * 
 * 输入：s = "abbcccddddeeeeedcba"
 * 输出：5
 * 解释：子字符串 "eeeee" 长度为 5 ，只包含字符 'e' 。
 * 
 * 
 * 示例 3：
 * 
 * 输入：s = "triplepillooooow"
 * 输出：5
 * 
 * 
 * 示例 4：
 * 
 * 输入：s = "hooraaaaaaaaaaay"
 * 输出：11
 * 
 * 
 * 示例 5：
 * 
 * 输入：s = "tourist"
 * 输出：1
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= s.length <= 500
 * s 只包含小写英文字母。
 * 
 * 
 */

// @lc code=start
class Solution1 {

    /**
     * @param String $s
     * @return Integer
     */
    // [1]for循环模拟
    // 执行用时：16 ms, 在所有 PHP 提交中击败了50.00%的用户
    // 内存消耗：15.2 MB, 在所有 PHP 提交中击败了62.50%的用户
    function maxPower($s) {
        $curlen= 1;  // 由于$s至少为1 那么就至少最大长度为1
        $max   = 1;  
        $pre   = ''; // 当前字符串
        $len   = strlen($s);
        for ($i=0; $i < $len; $i++) { 
            if ($s[$i] == $pre ) {
                // 如果当前字符等于前一个字符
                $curlen++; // 累加
                $max = max($max, $curlen);
            } else {
                // 如果当前字符等于前一个字符
                $max = max($max, $curlen);
                $pre = $s[$i]; // 当前字符
                $curlen = 1;    // 重置
            }
        }
        return $max;

    }
}

class Solution2 {

    /**
     * @param String $s
     * @return Integer
     */
    // [2]双指针
    // 执行用时：16 ms, 在所有 PHP 提交中击败了50.00%的用户
    // 内存消耗：15.3 MB, 在所有 PHP 提交中击败了12.50%的用户
    function maxPower($s) {
        $max = 0;
        $len = strlen($s);
        for ($i= 0; $i < $len; $i++) { 
            $j = $i;
            // 让 j保持不动，i继续向后移动
            while($s[$j] == $s[$i]) {
                $i++;
            }
            // 这里 $s[i] != $s[j] 两个字符不相等
            // 然后比较当前max和 i~j 哪个更大
            $max = max($max, $i - $j);
            $i--;// 然后向后移动一下i，从字符变化的地方开始
        }
        return $max;
    }
}
// @lc code=end

/*
【题目】
    给一个字符串，求寻找连续长度的字符串的最大长度。
    

【解题】
    1、for循环模拟
        max    初始值为1 （因为字符串长度默认至少为1，那么连续字符串至少也是1）
        curlen 表示当前扫描字符连续个数，初始值设置为1

        for => 循环 逐位扫描
            if  => 如果当前字符与前一个字符相等 s[i] == pre
                那么curlen++ (表示 连续字符增加)
                每次比较 max , curlen 哪个更大  (每次遍历一个字符都要更新 max)
            else =>如果当前字符与前一个字符不相等 s[i] != pre
                首先更新一下是否前一个字符的长度curlen 与当前最大值max （更新 max）
                然后重新赋值 pre = s[i]
                重新计算curlen = 1
        最后返回 max
        
        时间复杂度 O(N) 。字符串长度为N，需要全部遍历。
        空间复杂度 O(1) 。只用几个变量

    2、双指针
        两层循环，固定一个j指针，移动另一个i指针
        
        时间复杂度 O(N) 。字符串长度为N，需要全部遍历。
        空间复杂度 O(1) 。只用几个变量
        
                    
*/