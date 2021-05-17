<?php
/*
 * @Descripttion: 周二的题目补一下
 * @Author: tacks321@qq.com
 * @Date: 2021-05-12 16:11:13
 * @LastEditTime: 2021-05-12 19:00:54
 */

/*
 * @lc app=leetcode.cn id=55 lang=php
 *
 * [55] 跳跃游戏
 *
 * https://leetcode-cn.com/problems/jump-game/description/
 *
 * algorithms
 * Medium (41.54%)
 * Likes:    1184
 * Dislikes: 0
 * Total Accepted:    233.9K
 * Total Submissions: 555.5K
 * Testcase Example:  '[2,3,1,1,4]'
 *
 * 给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
 * 
 * 数组中的每个元素代表你在该位置可以跳跃的最大长度。
 * 
 * 判断你是否能够到达最后一个下标。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [2,3,1,1,4]
 * 输出：true
 * 解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [3,2,1,0,4]
 * 输出：false
 * 解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * 0 
 * 
 * 
 */

// @lc code=start
class Solution1 {

    /**
     * @param Integer[] $nums
     * @return Boolean
     */
    // 贪心
    // 执行用时：20 ms, 在所有 PHP 提交中击败了98.97%的用户
    // 内存消耗：18.4 MB, 在所有 PHP 提交中击败了82.47%的用户
    function canJump($nums) {
        $size = count($nums);
        $flag = 0; // 标记最远到达的位置
        for ($i=0; $i < $size ; $i++) { 
            if($i <= $flag) {
                $flag = max($flag, $i+$nums[$i]); // 最远可以到达的位置
                if($flag >=  $size - 1) {
                    return true;
                }
            }
        }
        return false;
    }

}

class Solution2 {

    /**
     * @param Integer[] $nums
     * @return Boolean
     */
    // 动规
    // 执行用时：32 ms, 在所有 PHP 提交中击败了26.80%的用户
    // 内存消耗：18.5 MB, 在所有 PHP 提交中击败了51.55%的用户
    function canJump($nums) {
        // dp[i] 表示 i是否可以到达
        $size= count($nums);
        $dp  = array_fill(0, $size, false);
        if($size > 0) {
            $dp[0] = true; // 第一个位置默认可达
        }
        // 遍历
        for ($i=1; $i < $size; $i++) { 
            // 想知道能否到达该点，则要看能否从该点之前的点跳跃过来，
            // 从后往前找的原因是越接近它的点越可能能够跳过来，可以减少循环次数
            for ($j=$i-1; $j >= 0 ; $j--) { 
                if($dp[$j] && $nums[$j] + $j >= $i) {
                    $dp[$i] = true;
                    break;
                }
            }
        }

        return $dp[$size-1];
    }

}
// @lc code=end

$obj = new Solution2();
$value= $obj->canJump([3,2,1,0,4]);
var_dump($value);


/*
【题目】
    问题是否能到达最后一个位置
        最后一个位置数字没有什么意义
    数组上每一个值，表示你当前可以最大跳的步数
        当这个数字足够大，那你可以一下跳到最后
    
【解析】

    贪心
        遍历整个数组
            每次更新最远可以跳到的位置，并标记 x + nums[x]
            如果 最远可以到达的位置 大于等于数组中的最后一个位置
                那就说明最后一个位置可达，我们就可以直接返回 true
        如果 遍历结束
            最后一个位置还不行，那么就返回 false

        时间复杂度：O(n) 
        空间复杂度：O(1) 

    动规
        这个题我们可以从后往前分析，首先判断倒数第二个元素能否到达最后一个元素，如果可以，我们将不再考虑最后一个元素，
        
        时间复杂度是O(n^2)
        空间复杂度是O(n)
*/