<?php
/*
 * @Descripttion: 周三的题
 * @Author: tacks321@qq.com
 * @Date: 2021-05-13 09:46:41
 * @LastEditTime: 2021-05-13 11:35:42
 */




/*
 * @lc app=leetcode.cn id=45 lang=php
 *
 * [45] 跳跃游戏 II
 *
 * https://leetcode-cn.com/problems/jump-game-ii/description/
 *
 * algorithms
 * Hard (38.27%)
 * Likes:    966
 * Dislikes: 0
 * Total Accepted:    129.2K
 * Total Submissions: 324.8K
 * Testcase Example:  '[2,3,1,1,4]'
 *
 * 给定一个非负整数数组，你最初位于数组的第一个位置。
 * 
 * 数组中的每个元素代表你在该位置可以跳跃的最大长度。
 * 
 * 你的目标是使用最少的跳跃次数到达数组的最后一个位置。
 * 
 * 假设你总是可以到达数组的最后一个位置。
 * 
 * 
 * 
 * 示例 1:
 * 
 * 
 * 输入: [2,3,1,1,4]
 * 输出: 2
 * 解释: 跳到最后一个位置的最小跳跃数是 2。
 * 从下标为 0 跳到下标为 1 的位置，跳 1 步，然后跳 3 步到达数组的最后一个位置。
 * 
 * 
 * 示例 2:
 * 
 * 
 * 输入: [2,3,0,1,4]
 * 输出: 2
 * 
 * 
 * 
 * 
 * 提示:
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
     * @return Integer
     */
    // 反向贪心
    // 执行用时：8 ms, 在所有 PHP 提交中击败了80.85%的用户
    // 内存消耗：15.2 MB, 在所有 PHP 提交中击败了61.70%的用户
    function jump($nums) {
        $position = count($nums) - 1;
        $steps = 0;

        while($position > 0) {
            for ($i=0; $i < $position; $i++) { 
                // 如果当前位置的可以跳跃
                if($i + $nums[$i] >= $position) {
                    // 找到最后一步跳跃前所在的位置之后
                    // 我们继续贪心地寻找倒数第二步跳跃前所在的位置
                    $position = $i; // 标记
                    
                    $steps++; // 步数++
                    // 因为是从小的下标遍历，所以如果这样贪心可以获取最小步数
                    break;
                }
            }
        }
        return $steps;
    }
}


class Solution2 {

    /**
     * @param Integer[] $nums
     * @return Integer
     */
    // 正向贪心
    // 执行用时：8 ms, 在所有 PHP 提交中击败了80.85%的用户
    // 内存消耗：15.1 MB, 在所有 PHP 提交中击败了91.49%的用户
    function jump($nums) {
        $size = count($nums) - 1; // 不访问最后一个位置
        $end  = 0; // 边界
        $maxpos= 0; // 最远跳的位置
        $steps = 0; // 步数
        
        for ($i=0; $i < $size; $i++) { 
            $maxpos = max($maxpos, $i + $nums[$i]); // 比一下当前哪个可以跳更远
            if($i >= $end) {
                //  目前能跳到的最远位置变成了下次起跳位置的有边界
                $end = $maxpos; // 更新最远边界
                $steps++; // 跳下一次
            }
        }
        return $steps;
    }
}
// @lc code=end


/*
【题目】
    跟昨天的《跳跃游戏Ⅰ》 有几个不同的地方

        题目已经确保 总是可以到达数组的最后一个位置。
        使用最少的跳跃次数到达数组的最后一个位置。
    
    贪心
        贪心算法，通过局部最优解得到全局最优解
    重点
        贪心策略的选取
【解析】
    反向贪心
        问题是找最少跳跃的次数，而且题目本来就说肯定能到最后一个位置
        那我们就假设我们现在已经在最后一个位置，我们向左去跳跃
        每次跳跃我们都尽可能(贪心)向左跳的远，也就是对应下标值更小，

        时间复杂度O(n*n)
        空间复杂度O(1)
    
    正向贪心
        「贪心」地进行正向查找，每次找到可到达的最远位置，就可以在线性时间内得到最少的跳跃次数。

        维护当前能够到达的最大下标位置，记为边界
        我们从左到右遍历数组，到达边界时，更新边界并将跳跃次数增加 1。
        

    
    
*/