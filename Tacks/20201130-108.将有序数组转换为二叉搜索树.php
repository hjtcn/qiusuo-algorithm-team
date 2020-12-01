<?php
/*
 * @Description:  新的一周 奥力给！！！
 * @Author: Tacks
 * @Date: 2020-11-30 17:48:33
 * @LastEditTime: 2020-11-30 18:30:45
 */
/*
 * @lc app=leetcode.cn id=108 lang=php
 *
 * [108] 将有序数组转换为二叉搜索树
 *
 * https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/description/
 *
 * algorithms
 * Easy (74.35%)
 * Likes:    645
 * Dislikes: 0
 * Total Accepted:    126.1K
 * Total Submissions: 169.4K
 * Testcase Example:  '[-10,-3,0,5,9]'
 *
 * 将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。
 * 
 * 本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。
 * 
 * 示例:
 * 
 * 给定有序数组: [-10,-3,0,5,9],
 * 
 * 一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：
 * 
 * ⁠     0
 * ⁠    / \
 * ⁠  -3   9
 * ⁠  /   /
 * ⁠-10  5
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
     * @param Integer[] $nums
     * @return TreeNode
     */
    // 执行用时：16 ms, 在所有 PHP 提交中击败了85.19%的用户
    // 内存消耗：18.6 MB, 在所有 PHP 提交中击败了35.85%的用户
    function sortedArrayToBST($nums) {
        // 执行结果仅代表一个可能的答案
        return $this->dfs($nums, 0, count($nums)-1);
    }

    function dfs($nums, $front, $rear) 
    {
        if($front > $rear) {
            return null;
        }
        // 中间节点作为根节点 ,根节点选取不同，结果不同
        $mid = floor(($front + $rear) / 2);
        $node = new TreeNode($nums[$mid]);
        $node->left = $this->dfs($nums, $front, $mid-1);
        $node->right = $this->dfs($nums, $mid + 1, $rear);
        return $node;
    }

}
// @lc code=end



/*

下面三个词的意思完全一样
    二叉搜索树（Binary Search Tree）
    有序二叉树（Ordered Binary Tree）
    排序二叉树（Sorted Binary Tree）
特点
    如果有左右子树，那么任意左节点的值都小于根节点，右节点都大于根节点
优势
    二叉查找树相比于其他数据结构的优势在于查找、插入的时间复杂度较低。为 O(logn)
构建
    中序遍历二叉查找树可得到一个关键字的有序序列
    一个无序序列可以通过构造一棵二叉查找树变成一个有序序列
    构造树的过程即为对无序序列进行查找的过程
    每次插入的新的结点都是二叉查找树上新的叶子结点，在进行插入操作时，不必移动其它结点，只需改动某个结点的指针，由空变为非空即可。
复杂度
    搜索、插入、删除的复杂度等于树高，期望 O(logn)，最坏 O(n)（数列有序，树退化成线性表）
拔高
    AVL 树、红黑树


【题目思考】
条件
    二叉搜索树的中序遍历，是一个升序的序列。
    二叉树的高度平衡，也就是有左右子树高度差不超过1
解决思路
    递归 DFS 深度优先遍历


*/