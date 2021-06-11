<?php
/*
 * @Descripttion:  周五端午节快乐
 * @Author: tacks321@qq.com
 * @Date: 2021-06-11 16:38:00
 * @LastEditTime: 2021-06-11 17:37:04
 */

/*
 * @lc app=leetcode.cn id=40 lang=php
 *
 * [40] 组合总和 II
 *
 * https://leetcode-cn.com/problems/combination-sum-ii/description/
 *
 * algorithms
 * Medium (63.59%)
 * Likes:    604
 * Dislikes: 0
 * Total Accepted:    164.6K
 * Total Submissions: 258.8K
 * Testcase Example:  '[10,1,2,7,6,1,5]\n8'
 *
 * 给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 * 
 * candidates 中的每个数字在每个组合中只能使用一次。
 * 
 * 说明：
 * 
 * 
 * 所有数字（包括目标数）都是正整数。
 * 解集不能包含重复的组合。 
 * 
 * 
 * 示例 1:
 * 
 * 输入: candidates = [10,1,2,7,6,1,5], target = 8,
 * 所求解集为:
 * [
 * ⁠ [1, 7],
 * ⁠ [1, 2, 5],
 * ⁠ [2, 6],
 * ⁠ [1, 1, 6]
 * ]
 * 
 * 
 * 示例 2:
 * 
 * 输入: candidates = [2,5,2,1,2], target = 5,
 * 所求解集为:
 * [
 * [1,2,2],
 * [5]
 * ]
 * 
 */

// @lc code=start
class Solution {

    protected $result = [];
    /**
     * @param Integer[] $candidates
     * @param Integer $target
     * @return Integer[][]
     */
    // 执行用时：8 ms, 在所有 PHP 提交中击败了100.00%的用户
    // 内存消耗：15.4 MB, 在所有 PHP 提交中击败了7.14%的用户
    function combinationSum2($candidates, $target) {
        if($target <= 0) {
            return [];
        }
        // 升序排序，用于剪枝，提高搜索速度，另外也是为去重
        sort($candidates);
        $this->_dfs($candidates,count($candidates), $target, [], 0);
        return $this->result;
    }

    /**
     * 回溯
     *
     * @param array $nums  元素数组
     * @param int $len     数组大小
     * @param int $target  目标值、下一个目标值..依次减小
     * @param array $path  组合
     * @param int $start   下一次选择搜索起点
     */
    private function _dfs($nums, $len, $target, $path, $start) {
        // target 为负数和 0 的时候 便不会向下产生新的节点
        if($target < 0 ) {
            return ; // 此路不通 
        }
        // 找到一个满足的组合
        if($target == 0) {
            $this->result[] = $path; // 加入结果集
            return ;
        }

        // 从 start 遍历 就是为了防止元素由于顺序上不同，而导致的组合重复
        for ($i=$start; $i< $len; $i++) { 
            // 剪枝1  当前值已经小于0，后面肯定也小于0 ，不必向下
            if($target - $nums[$i] < 0) {
                break;
            }
            // 【重点】剪枝2  i>start 就是在不同的层级上的 
            if( ($i > $start) && ($nums[$i] == $nums[$i-1]) ) {
                continue;
            }
            // 加入组合中
            $path[] = $nums[$i];

            // 【重点】选取下一个target 数字不能重复使用 递归向下，要++1
            $this->_dfs($nums, $len, $target-$nums[$i], $path, $i+1);
            // 回溯
            array_pop($path);
        }
    }
    
}
// @lc code=end

/*
【题目】
组合总和
    第 39 题
    第 40 题
不同点是
    candidates 中的数字可以无限制重复被选取；
    candidates 中的每个数字在每个组合中只能使用一次。
相同点是
    相同数字列表的不同排列视为一个结果。

    那么解法，可以直接套用39题思路，然后加上去重即可

    剪枝：去重思路
        1、哈希
        2、顺序排列 判断重复则跳过
    

【解析】    
    https://leetcode-cn.com/problems/combination-sum-ii/solution/hui-su-suan-fa-jian-zhi-python-dai-ma-java-dai-m-3/
    DFS是一种思想，可以形成一种模板的
    
        dfs(路径, 选择列表, start遍历位置):
            if 满足结束条件:
                跳过
                return
            if 满足结束条件:
                add(路径)
                return

            for start遍历的位置
                剪枝

                push路径
                
                dfs(路径, 选择列表, start+1 / start )

                pop(路径)

 

    重点在于剪枝的过程


    candidates = [2,5,2,1,2], target = 5

    树的形成
        target作为树的根节点，然后依次减去数组中的值，向下深度优先遍历，找到下面的子节点

        5 作为根节点
        sort(candidates)   [1 2 2 2 5]  
        然后向下形成子节点 [4 3 3 3 0] 由于同一层如果减去上一层的数字一样只需要保留一个分支即可 [4 3 0]

    避免重复的思想
                  1
                 / \
                2   2  这种情况不会发生 但是却允许了不同层级之间的重复即：
               /     \
              5       5
                例2
                  1
                 /
                2      这种情况确是允许的
               /
              2  

    
*/