<?php

// 周二 20201110
/*

昨天晚上双十一，真的是剁手哦，呜呜呜，有史以来是我花最多的一次

京东 淘宝 都冲了

*/

/*
 * @lc app=leetcode.cn id=104 lang=php
 *
 * [104] 二叉树的最大深度
 *
 * https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (75.25%)
 * Likes:    734
 * Dislikes: 0
 * Total Accepted:    301.5K
 * Total Submissions: 400.4K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，找出其最大深度。
 * 
 * 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
 * 
 * 说明: 叶子节点是指没有子节点的节点。
 * 
 * 示例：
 * 给定二叉树 [3,9,20,null,null,15,7]，
 * 
 * ⁠   3
 * ⁠  / \
 * ⁠ 9  20
 * ⁠   /  \
 * ⁠  15   7
 * 
 * 返回它的最大深度 3 。
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
     * @return Integer
     */
    // 递归
    // 【深度优先搜索】 DFS
    // 执行用时：16 ms, 在所有 PHP 提交中击败了45.38%的用户
    // 内存消耗：17 MB, 在所有 PHP 提交中击败了17.63%的用户
    function maxDepth1($root) {
        if(!$root) return 0;
        // 递归真香
        $left  = $this->maxDepth1($root->left);
        $right = $this->maxDepth1($root->right);
        // 如果一个二叉树，知道左子树和右子树的最大深度 l 和 r，那么该二叉树的最大深度即为 max(l,r)+1
        return max($left, $right) + 1;
    }
    

    // 【广度优先遍历】BFS 
    // 执行用时：8 ms, 在所有 PHP 提交中击败了97.59%的用户
    // 内存消耗：16.9 MB, 在所有 PHP 提交中击败了19.26%的用户
    function maxDepth($root) {
        if(!$root) {
            return 0;
        }
        $depth = 0; 
        $nodeList = [$root];
        $count = 1; // 当前层的节点数
        $count_next = 0; // 统计下一层push进去的节点数
        while($nodeList) {
            $tree = array_shift($nodeList); // 弹出
            $count--;
            if($tree->left) {
                array_push($nodeList, $tree->left); // 左子树
                $count_next++;
            }
            if($tree->right) {
                array_push($nodeList, $tree->right); 
                $count_next++;
            }
            // 如果当前层的节点数目遍历结束
            if( $count == 0 ){
                $depth++;
                $count = $count_next; // 开始遍历下一层
                $count_next = 0;
            }

        }
        return $depth;
    }


}
// @lc code=end


/*

【深度优先搜索】
    递归
    如果一个二叉树，知道左子树和右子树的最大深度 l 和 r，那么该二叉树的最大深度即为 max(l,r)+1

    同样的
    左子树和右子树的最大深度又可以以同样的方式进行计算。可以先递归计算出其左子树和右子树的最大深度

    时间复杂度
        O(n)，其中 nn 为二叉树节点的个数。每个节点在递归中只被遍历一次
    空间复杂度
        O(height)，其中  height 表示二叉树的高度。
        递归函数需要栈空间，而栈空间取决于递归的深度，因此空间复杂度等价于二叉树的高度
 
【广度优先遍历】BFS 
    本题，需要广度优先搜索的队列里存放的是「当前层的所有节点」
    一层一层遍历，所以也叫层级遍历，
    变量 ans 来维护拓展的次数，每次拓展的是本层所有的节点
*/