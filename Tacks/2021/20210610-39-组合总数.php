<?php
/*
 * @Descripttion: 周四的题，好久没刷了
 * @Author: tacks321@qq.com
 * @Date: 2021-06-10 17:57:07
 * @LastEditTime: 2021-06-10 19:20:26
 */

/*
 * @lc app=leetcode.cn id=39 lang=php
 *
 * [39] 组合总和
 *
 * https://leetcode-cn.com/problems/combination-sum/description/
 *
 * algorithms
 * Medium (72.47%)
 * Likes:    1383
 * Dislikes: 0
 * Total Accepted:    269.3K
 * Total Submissions: 371.6K
 * Testcase Example:  '[2,3,6,7]\n7'
 *
 * 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
 * 
 * candidates 中的数字可以无限制重复被选取。
 * 
 * 说明：
 * 
 * 
 * 所有数字（包括 target）都是正整数。
 * 解集不能包含重复的组合。 
 * 
 * 
 * 示例 1：
 * 
 * 输入：candidates = [2,3,6,7], target = 7,
 * 所求解集为：
 * [
 * ⁠ [7],
 * ⁠ [2,2,3]
 * ]
 * 
 * 
 * 示例 2：
 * 
 * 输入：candidates = [2,3,5], target = 8,
 * 所求解集为：
 * [
 * [2,2,2,2],
 * [2,3,3],
 * [3,5]
 * ]
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= candidates.length <= 30
 * 1 <= candidates[i] <= 200
 * candidate 中的每个元素都是独一无二的。
 * 1 <= target <= 500
 * 
 * 
 */

// @lc code=start
class Solution {

    protected $result = []; // 结果组合
    /**
     * @param Integer[] $candidates
     * @param Integer $target
     * @return Integer[][]
     */
    // 执行用时：12 ms, 在所有 PHP 提交中击败了100.00%的用户
    // 内存消耗：15.3 MB, 在所有 PHP 提交中击败了72.00%的用户
    function combinationSum($candidates, $target) {
        if($target <= 0) {
            return [];
        }
        // 排序，用于剪枝，提高搜索速度
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
            // 剪枝
            if($target - $nums[$i] < 0) {
                break;
            }
            // 加入组合中
            $path[] = $nums[$i];
            // 选取下一个target 数字可以重复使用
            $this->_dfs($nums, $len, $target-$nums[$i], $path, $i);
            // 回溯
            array_pop($path);
        }
    }
}
// @lc code=end


/*
【题意】
    就是要在一个元素数组中，找几个值的组合，相加之和为target目标值，然后要找到所有的组合，返回结果集。

    组合中，元素可以重复使用。
    
【解析】

    回溯法
        可以想象成一个决策树，每个节点都要进行遍历，并进行决策
            1、路径：当前已经做出来选择的 path
            2、选择列表： 当前可以选择的列表 nums start来控制选择列表范围
            3、结束条件：当前决策树达到叶子节点
        关于组合问题的
            例如 [a,b] 与 [b,a] 在本题中是一种结果，所以可以利用start来进行选择 数组元素的范围
        关于排列问题的
            例如 [a,b] 与 [b,a] 如果是两种结果，则需要用 used数组来标记
        

        递归：深度优先递归DFS
            从根节点出发，搜索到整个树
        剪枝：对于元素集合进行sort排序
            是为了让相同的元素在一起，便于剪枝
            如果当前元素 candidates[i] > target 那么不需要向下遍历了，直接回溯到上一层
        回溯：
            搜索到任意一个节点，会进行决策判断，如果不符合要求，便进行向父节点回溯，回到之前的状态



    时间复杂度：O(S) S为targetArr的长度。实际运行情况远远小于(N*2^N) 
    空间复杂度：O(target)
*/
