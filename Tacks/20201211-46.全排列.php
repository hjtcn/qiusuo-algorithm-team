<?php
/*
 * @Descripttion: 周五开心，周末冲
 * @Author: tacks321@qq.com
 * @Date: 2020-12-11 10:26:41
 * @LastEditTime: 2020-12-11 11:11:49
 */


/*
 * @lc app=leetcode.cn id=46 lang=php
 *
 * [46] 全排列
 *
 * https://leetcode-cn.com/problems/permutations/description/
 *
 * algorithms
 * Medium (77.08%)
 * Likes:    1029
 * Dislikes: 0
 * Total Accepted:    227.9K
 * Total Submissions: 295K
 * Testcase Example:  '[1,2,3]'
 *
 * 给定一个 没有重复 数字的序列，返回其所有可能的全排列。
 * 
 * 示例:
 * 
 * 输入: [1,2,3]
 * 输出:
 * [
 * ⁠ [1,2,3],
 * ⁠ [1,3,2],
 * ⁠ [2,1,3],
 * ⁠ [2,3,1],
 * ⁠ [3,1,2],
 * ⁠ [3,2,1]
 * ]
 * 
 */

// @lc code=start
class Solution {

    public $res = [];

    /**
     * @param Integer[] $nums
     * @return Integer[][]
     */
    function permute($nums) {
        $this->_backtracking($nums, []);
        return $this->res;
    }

    /**
     * 回溯
     * @param [array] $nums 原始数组nums [1,2,3]
     * @param [array] $arr  当前数组arr  []
     */
    private function _backtracking($nums, $arr) {
        // 当前arr的个数已经等于原始数组nums
        if(count($arr) == count($nums)) {
            $this->res[] = $arr;
            return ;
        }else{
            // 循环遍历原始数组nums,当前值不存在arr中时，将该值放入arr中
            foreach($nums as $value) {
                if( !in_array($value, array($arr) ) ) {
                    // 每次放一个值，剩下的进行全排列
                    $arr[] = $value;
                    $this->_backtracking($nums, $arr);
                    // 将arr进行出栈操作，继续遍历
                    array_pop($arr);
                }
            }
        }
        
    }

}
// @lc code=end



/*
【回溯法】

    一种通过探索所有可能的候选解来找出所有的解的算法。
    如果候选解被确认不是一个解的话（或者至少不是最后一个解），回溯算法会通过在上一步进行一些变化抛弃该解，即回溯并且再次尝试。


回溯看起来有点懵逼啊



*/