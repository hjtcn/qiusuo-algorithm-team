<?php
/*
 * @Descripttion: 周四了
 * @Author: tacks321@qq.com
 * @Date: 2021-06-17 17:31:32
 * @LastEditTime: 2021-06-17 18:25:06
 */

/*
 * @lc app=leetcode.cn id=47 lang=php
 *
 * [47] 全排列 II
 *
 * https://leetcode-cn.com/problems/permutations-ii/description/
 *
 * algorithms
 * Medium (63.30%)
 * Likes:    725
 * Dislikes: 0
 * Total Accepted:    175.6K
 * Total Submissions: 277.3K
 * Testcase Example:  '[1,1,2]'
 *
 * 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,1,2]
 * 输出：
 * [[1,1,2],
 * ⁠[1,2,1],
 * ⁠[2,1,1]]
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [1,2,3]
 * 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 
 * -10 
 * 
 * 
 */

// @lc code=start
class Solution {

    // 结果集
    protected $result = [];

    /**
     * @param Integer[] $nums
     * @return Integer[][]
     */
    // 执行用时：24 ms, 在所有 PHP 提交中击败了44.26%的用户
    // 内存消耗：15.8 MB, 在所有 PHP 提交中击败了44.26%的用户
    function permuteUnique($nums) {
        $size = count($nums);
        // 边界判断
        if($size == 0) {
            return $this->result;
        }
        // 升序 用于剪枝 去重
        sort($nums);

        $used = [];
        $path = [];
        $depth= 0; // 当前层
        $this->_dfs($nums, $size, $depth, $path, $used);

        return $this->result;
    }
 

    /**
     * DFS 深搜回溯
     *
     * @param array $nums  元素数组
     * @param int $size    数组大小
     * @param int $depth   当前层
     * @param array $path  排列
     * @param array $used  去重
     */
    private function _dfs($nums, $size, $depth, $path, $used) {
        // 找到一个满足的排列
        if($depth == $size) {
            $this->result[] = $path; // 加入结果集
            return ;
        }

        for ($i=0; $i< $size; $i++) { 
            // 【重点】剪枝
            // 使用过了就不再使用，可以去重一部分
            if($used[$i]) {
                continue;
            }
            
            // 【重点】剪枝 
            //  与同一层的 $nums[$i-1] 刚刚被撤销，正是因为刚被撤销，下面的搜索中还会使用到，因此会产生重复，剪掉的应该这样的分支
            if( ($i > 0) && ($nums[$i] == $nums[$i-1]) && !$used[$i-1] ) {
                continue;
            }
            // 加入排列中
            $path[] = $nums[$i];
            
            $used[$i] = true; // 标记

            // 递归处理
            $this->_dfs($nums, $size, $depth+1, $path, $used);

            // 回溯
            $used[$i] = false;
            array_pop($path);
        }
    }
}


// @lc code=end

/*
【题目】
全排列类型
    [46] 全排列 https://leetcode-cn.com/problems/permutations/description/
        给定一个 没有重复 数字的序列，返回其所有可能的全排列。
    [47] 全排列 II https://leetcode-cn.com/problems/permutations-ii/description/
        给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。

    区别是：两组数列是否有重复


组合类型
    [39]组合总和 (https://leetcode-cn.com/problems/combination-sum/)
    [40]组合总和II (https://leetcode-cn.com/problems/combination-sum-ii/)

    区别是：两组数列是否可以选择相同的数

    关于排列问题的
            例如 [a,b] 与 [b,a] 如果是两种结果，则需要用 used 数组来标记去重

    关于组合问题的
            例如 [a,b] 与 [b,a] 如果是一种结果，可以根据顺序按照 start 进行去重
 

全排列（枚举所有解）
    回溯法
        回溯就是类似枚举的搜索尝试过程，主要是在搜索尝试过程中寻找问题的解，当发现已不满足求解条件时，就“回溯”返回，尝试别的路径。
        枚举的搜索尝试，这里指DFS进行深搜


【解析】

    DFS+回溯


    ？？待细化总结
    

*/