<?php
/*
 * @Descripttion: 【每日一题】加油，虽然今天明显有点忙，但是还是要坚持干
 * @Author: tacks321@qq.com
 * @Date: 2021-01-05 20:41:55
 * @LastEditTime: 2021-01-06 08:53:41
 */


/*
 * @lc app=leetcode.cn id=830 lang=php
 *
 * [830] 较大分组的位置
 *
 * https://leetcode-cn.com/problems/positions-of-large-groups/description/
 *
 * algorithms
 * Easy (47.81%)
 * Likes:    93
 * Dislikes: 0
 * Total Accepted:    39.9K
 * Total Submissions: 73.8K
 * Testcase Example:  '"abbxxxxzzy"'
 *
 * 在一个由小写字母构成的字符串 s 中，包含由一些连续的相同字符所构成的分组。
 * 
 * 例如，在字符串 s = "abbxxxxzyy" 中，就含有 "a", "bb", "xxxx", "z" 和 "yy" 这样的一些分组。
 * 
 * 分组可以用区间 [start, end] 表示，其中 start 和 end 分别表示该分组的起始和终止位置的下标。上例中的 "xxxx"
 * 分组用区间表示为 [3,6] 。
 * 
 * 我们称所有包含大于或等于三个连续字符的分组为 较大分组 。
 * 
 * 找到每一个 较大分组 的区间，按起始位置下标递增顺序排序后，返回结果。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：s = "abbxxxxzzy"
 * 输出：[[3,6]]
 * 解释："xxxx" 是一个起始于 3 且终止于 6 的较大分组。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：s = "abc"
 * 输出：[]
 * 解释："a","b" 和 "c" 均不是符合要求的较大分组。
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：s = "abcdddeeeeaabbbcd"
 * 输出：[[3,5],[6,9],[12,14]]
 * 解释：较大分组为 "ddd", "eeee" 和 "bbb"
 * 
 * 示例 4：
 * 
 * 
 * 输入：s = "aba"
 * 输出：[]
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * s 仅含小写英文字母
 * 
 * 
 */

// @lc code=start
class Solution {

    /**
     * @param String $s
     * @return Integer[][]
     */
    // 执行用时：12 ms, 在所有 PHP 提交中击败了20.00%的用户
    // 内存消耗：15.3 MB, 在所有 PHP 提交中击败了50.00%的用户
    // 一次遍历
    function largeGroupPositions($s) {
        $result = []; // 结果数组
        $len = strlen($s);
        if($len <= 2) {
            return $result;
        }

        $start = $end = 0;
        
        for($i=1; $i<$len; $i++) {
            // 如果前后两个值不相等
            if($s[$i] !== $s[$i-1]) {
                // 判断两个指针的区间范围
                if($end - $start >= 2) {
                    $result[] = [$start, $end]; // 较大分组 符合要求加入结果数组中
                }
                $start = $end = $i; // 双指针重置为当前位置
            }else{
                // 如果当前值与前面的一样，就移动end指针
                $end++;
            }
        }
        // 最后在判断一次
        if($end - $start >= 2) {
            $result[] = [$start, $end];
        }
        return $result;
    }
}
// @lc code=end

/*

较大分组：大于或等于三个连续字符的分组（双指针的使用，用来确定较大分组的区间）


时间复杂度：O(n)，其中 n 是字符串的长度。只需要遍历一次该数组。
空间复杂度：O(1)。我们只需要常数的空间来保存若干变量，返回值不计入空间复杂度。

 
*/