<?php
/*
 * @Descripttion: 补一下周三的题
 * @Date: 2020-12-03 22:04:02
 * @LastEditTime: 2021-11-08 09:57:35
 */

/*
 * @lc app=leetcode.cn id=257 lang=php
 *
 * [257] 二叉树的所有路径
 *
 * https://leetcode-cn.com/problems/binary-tree-paths/description/
 *
 * algorithms
 * Easy (66.17%)
 * Likes:    407
 * Dislikes: 0
 * Total Accepted:    86.5K
 * Total Submissions: 130.5K
 * Testcase Example:  '[1,2,3,null,5]'
 *
 * 给定一个二叉树，返回所有从根节点到叶子节点的路径。
 * 
 * 说明: 叶子节点是指没有子节点的节点。
 * 
 * 示例:
 * 
 * 输入:
 * 
 * ⁠  1
 * ⁠/   \
 * 2     3
 * ⁠\
 * ⁠ 5
 * 
 * 输出: ["1->2->5", "1->3"]
 * 
 * 解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3
 * 
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     public $val = null;
 *     public $left = null;
 *     public $right = null;
 *     function __construct($value) { $this->val = $value; }
 * }
 */
class Solution {

    protected $ans = [];

    /**
     * @param TreeNode $root
     * @return String[]
     */
    function binaryTreePaths($root) {
        $this->dfs($root, '');
        return $this->ans;
    }

    // 执行用时：0 ms, 在所有 PHP 提交中击败了100.00%的用户
    // 内存消耗：15.1 MB 在所有 PHP 提交中击败了8.13%的用户
    private function dfs($root, $path)
    {
        // 空节点
        if($root === null) return ;
        // 拼接路径
        $path .= ($path ? '->' : '') . $root->val;
        // 遍历到叶子节点 终止
        if($root->left === null && $root->right === null) {
            // 如果当前节点是叶子节点，则在当前路径末尾添加该节点后我们就得到了一条从根节点到叶子节点的路径
            $this->ans[] = $path;
        }
        // 如果当前节点不是叶子节点，则在当前的路径末尾添加该节点，并继续递归遍历该节点的每一个孩子节点。
        // 左
        $this->dfs($root->left, $path);
        // 右
        $this->dfs($root->right, $path);
        
    }

}
// @lc code=end

/*
遍历二叉树 DFS

最直观的方法是使用深度优先搜索。在深度优先搜索遍历二叉树时，我们需要考虑当前的节点以及它的孩子节点。


*/