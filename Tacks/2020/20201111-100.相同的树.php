<?php

// 周一 20201109

/*
 * @lc app=leetcode.cn id=100 lang=php
 *
 * [100] 相同的树
 *
 * https://leetcode-cn.com/problems/same-tree/description/
 *
 * algorithms
 * Easy (60.20%)
 * Likes:    505
 * Dislikes: 0
 * Total Accepted:    153.5K
 * Total Submissions: 254.9K
 * Testcase Example:  '[1,2,3]\n[1,2,3]'
 *
 * 给定两个二叉树，编写一个函数来检验它们是否相同。
 * 
 * 如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
 * 
 * 示例 1:
 * 
 * 输入:       1         1
 * ⁠         / \       / \
 * ⁠        2   3     2   3
 * 
 * ⁠       [1,2,3],   [1,2,3]
 * 
 * 输出: true
 * 
 * 示例 2:
 * 
 * 输入:      1          1
 * ⁠         /           \
 * ⁠        2             2
 * 
 * ⁠       [1,2],     [1,null,2]
 * 
 * 输出: false
 * 
 * 
 * 示例 3:
 * 
 * 输入:       1         1
 * ⁠         / \       / \
 * ⁠        2   1     1   2
 * 
 * ⁠       [1,2,1],   [1,1,2]
 * 
 * 输出: false
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
 *     function __construct($val = 0, $left = null, $right = null) {
 *         $this->val = $val;
 *         $this->left = $left;
 *         $this->right = $right;
 *     }
 * }
 */
class Solution {

    /**
     * @param TreeNode $p
     * @param TreeNode $q
     * @return Boolean
     */
    // 判断对象是否相同即可
    // 执行用时：12 ms, 在所有 PHP 提交中击败了11.81%的用户
    // 内存消耗：15.3 MB, 在所有 PHP 提交中击败了5.12%的用户
    function isSameTree1($p, $q) {
        return $p == $q;
    }

    // 深搜
    // 执行用时：8 ms, 在所有 PHP 提交中击败了69.29%的用户
    // 内存消耗：15.3 MB, 在所有 PHP 提交中击败了5.12%的用户
    function isSameTree2($p, $q) {
        // 两个二叉树都为空
        if($p == null && $q == null ) {
            return true;
        }elseif( $p == null || $q == null) {
            // 其中一边走到最后，两边长度不等。 两个二叉树中有且只有一个为空，则两个二叉树一定不相同
            return false;
        }elseif( $p->val != $q->val) {
            // 对应节点位置的值 不相等。根节点的值，若不相同则两个二叉树一定不同
            return false;
        }
        // 若根节点相同，再分别判断两个二叉树的左子树是否相同以及右子树是否相同
        return $this->isSameTree2($p->left, $q->left) && $this->isSameTree2($p->right, $q->right);
    }

    // 广搜
    // 执行用时：8 ms, 在所有 PHP 提交中击败了69.29%的用户
    // 内存消耗：15.3 MB, 在所有 PHP 提交中击败了5.12%的用户
    function isSameTree3($p, $q) {
        // 两个队列同时为空，则两个二叉树相同
        if($p == null && $q == null) {
            return true;
        }elseif($p == null || $q == null) {
            return false;
        }

        // 两个二叉树的根节点分别加入两个队列
        $arr_p = [$p];
        $arr_q = [$q];
        while($arr_p && $arr_q) {
            // 弹出队列
            $node_p = array_shift($arr_p);
            $node_q = array_shift($arr_q);
            // 如果两个节点的值不相同则两个二叉树一定不同
            if($node_p->val != $node_q->val) {
                return false;
            }

            // 如果只有一个节点的左子节点为空，或者只有一个节点的右子节点为空，则两个二叉树的结构不同，因此两个二叉树一定不同
            // 左子树不一样，有一个为空
            if ( ($node_p->left == null ) ^ ($node_q->left == null)) {
                return false;
            }
            // 右子树不一样，有一个为空
            if ( ($node_p->right==null) ^ ($node_q->right== null) ) {
                return false;
            }
           
            // 如果两个节点的子节点的结构相同，则将两个节点的非空子节点分别加入两个队列
            // 子节点加入队列时需要注意顺序，如果左右子节点都不为空，则先加入左子节点，后加入右子节点。
            if($node_p->left) {
                array_push($arr_p, $node_p->left);// 先加入左子节点
            }
            if($node_p->right) {
                array_push($arr_p, $node_p->right);// 后加入右子节点
            }
            if($node_q->left) {
                array_push($arr_q, $node_q->left);
            }
            if($node_q->right) {
                array_push($arr_q, $node_q->right);
            }

        }

       return empty($arr_p) && empty($arr_q);
 
    }
}
// @lc code=end


/*

比较两个树是否相等，那么就是所有对应节点的值相同

对象
    这里的树节点采用的是对象进行模拟的，直接比较两个对象是否相等

深搜
    深度优先搜索
        递归的过程！
        两个二叉树都为空，则两个二叉树相同
        两个二叉树中有且只有一个为空，则两个二叉树一定不相同
        两个二叉树都不为空，那么首先判断它们的根节点的值是否相同，若不相同则两个二叉树一定不同
        若根节点相同，再分别判断两个二叉树的左子树是否相同以及右子树是否相同

广搜
    广度优先搜索
        同样首先判断两个二叉树是否为空，如果两个二叉树都不为空，则从两个二叉树的根节点开始广度优先搜索

        迭代遍历的过程！
        初始时将两个二叉树的根节点分别加入两个队列
        每次从两个队列各取出一个节点，进行如下比较操作
            比较两个节点的值，如果两个节点的值不相同则两个二叉树一定不同
            如果两个节点的值相同，则判断两个节点的子节点是否为空，如果只有一个节点左子树或者右子树为空，那么两个二叉树结构不同
            如果两个节点的子节点的结构相同
                则将两个节点的非空子节点分别加入两个队列，子节点加入队列时需要注意顺序，如果左右子节点都不为空，则先加入左子节点，后加入右子节点


*/