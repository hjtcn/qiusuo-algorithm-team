<?php

// 周三 20201111


/*

堆积了好几天的题，今天抽空一起看了吧，懒惰

*/

/*
 * @lc app=leetcode.cn id=107 lang=php
 *
 * [107] 二叉树的层次遍历 II
 *
 * https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/description/
 *
 * algorithms
 * Easy (67.67%)
 * Likes:    364
 * Dislikes: 0
 * Total Accepted:    108.6K
 * Total Submissions: 160.5K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
 * 
 * 例如：
 * 给定二叉树 [3,9,20,null,null,15,7],
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 
 * 
 * 返回其自底向上的层次遍历为：
 * 
 * [
 * ⁠ [15,7],
 * ⁠ [9,20],
 * ⁠ [3]
 * ]
 * 
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

    /**
     * @param TreeNode $root
     * @return Integer[][]
     */
    // 执行用时：4 ms, 在所有 PHP 提交中击败了100.00%的用户
    // 内存消耗：16 MB, 在所有 PHP 提交中击败了6.25%的用户
    function levelOrderBottom($root) {
        if(!$root) {
            return [];
        }

        $arr = $res = [];
        $level = 0;
        $level++;
        array_push($arr, $root);

        while($count = count($arr)) {
            for($i=0; $i<$count; $i++) {
                $node = array_shift($arr);
                $res[$level][] = $node->val;
                if($node->left) {
                    array_push($arr, $node->left);
                }
                if($node->right) {
                    array_push($arr, $node->right);
                }
            }
            $level++;
        }
        return array_reverse($res);
       
    }
}
// @lc code=end

/*

其实跟这个题类似，「102. 二叉树的层序遍历」
    https://leetcode-cn.com/problems/binary-tree-level-order-traversal/1/
只是要求最后二叉树层次遍历输出的结果正好相反


【广度优先搜索】BFS
从根节点开始搜索，每次遍历同一层的全部节点，使用一个列表存储该层的节点值。

如果要求从上到下输出每一层的节点值，做法是很直观的，
在遍历完一层节点之后，将存储该层节点值的列表添加到结果列表的尾部。

本题最后反转了一下
 
*/