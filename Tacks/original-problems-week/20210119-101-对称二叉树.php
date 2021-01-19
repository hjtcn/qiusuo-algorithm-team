<?php
/*
 * @Descripttion: 周二 奥里给！！！ 感觉一天天好快啊 
 * @Author: tacks321@qq.com
 * @Date: 2021-01-19 17:20:25
 * @LastEditTime: 2021-01-19 18:21:39
 */

/*
 * @lc app=leetcode.cn id=101 lang=php
 *
 * [101] 对称二叉树
 *
 * https://leetcode-cn.com/problems/symmetric-tree/description/
 *
 * algorithms
 * Easy (53.37%)
 * Likes:    1203
 * Dislikes: 0
 * Total Accepted:    254K
 * Total Submissions: 475.8K
 * Testcase Example:  '[1,2,2,3,4,4,3]'
 *
 * 给定一个二叉树，检查它是否是镜像对称的。
 * 
 * 
 * 
 * 例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
 * 
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠/ \ / \
 * 3  4 4  3
 * 
 * 
 * 
 * 
 * 但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
 * 
 * ⁠   1
 * ⁠  / \
 * ⁠ 2   2
 * ⁠  \   \
 * ⁠  3    3
 * 
 * 
 * 
 * 
 * 进阶：
 * 
 * 你可以运用递归和迭代两种方法解决这个问题吗？
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
class Solution1 {

    /**
     * @param TreeNode $root
     * @return Boolean
     */
    // 递归 时间复杂度 O(n)  空间复杂度 O(n)
    // 执行用时：8 ms, 在所有 PHP 提交中击败了77.48%的用户
    // 内存消耗：15.4 MB, 在所有 PHP 提交中击败了45.05%的用户
    function isSymmetric($root) {
        return $this->_check($root, $root);
    }

    // 递归函数
    function _check($p, $q) {
        // 边界判断
        if($p == null && $q == null) {
            return true;
        }

        if($p == null || $q == null) {
            return false;
        }
        
        // 首先判断当前值是否相等
        // 递归处理左右子树
        return ($p->val == $q->val)
                && $this->_check($p->left, $q->right)
                && $this->_check($p->right, $q->left);
    }
}
// @lc code=end


 
class Solution2 {

    /**
     * @param TreeNode $root
     * @return Boolean
     */
    // 迭代 时间复杂度 O(n)  空间复杂度 O(n)
    // 执行用时：12 ms, 在所有 PHP 提交中击败了39.64%的用户
    // 内存消耗：15.3 MB, 在所有 PHP 提交中击败了86.49%的用户
    function isSymmetric($root) {
        $stack = []; // 引入栈
        // 初始化：PHP array_push 可以一次向stack压入多个值
        array_push($stack, $root, $root);
        // 直到栈为空
        while(!empty($stack)) {
            // 弹出栈顶 一次出栈两个进行比较
            $p = array_pop($stack);
            $q = array_pop($stack);
            // 边界判断
            if($p == null && $q == null) {
                continue;
            }
            if($p == null || $q == null || $p->val != $q->val) {
                return false;
            }
            // 如果左右子树的根节点值相等，则将左子树的 left 、右子树的 right； 左子树的 right 、右子树的 left 依次入栈
            array_push($stack, $p->left, $q->right);
            array_push($stack, $p->right, $q->left);
            // 继续向下遍历
        }
        return true;
    }

    
}

  
    
 
/*

[题目]
    给定一个二叉树，检查它是否是镜像对称的。

    问题：什么是镜像对称？  => 如果一个树的左子树与右子树镜像对称，那么这个树是对称的。
    问题：镜像的对称条件？  => 两个子树有相同的根节点，而且每颗的子树都互相对应相等。

【递归】
    递归的 同步移动两个指针来遍历左子树和右子树
    抽象出来就是一个二叉树A，它的根节点是A，他左右两个子树分别是 p 、 q 

    确定递归的三要素
        (a)确定递归函数的参数和返回值
        (b)确定终止条件
        (c)确定单层递归的逻辑

    那么首先边界判断
            判断 p q 是否都为空   => 对称
            如果 p q 其中一个为空 => 不对称
    然后判断 当前节点 p 和 q 的值是否相等  => 相等则继续递归左右子树
        由于是二叉树，所以这里对应两个指针，分别去移动位置，然后判断当前节点值是否相等即可。
            p向左移动 q向右移动
            p向右移动 q向左移动


    
    时间复杂度：
        这里遍历了这棵树，时间复杂度为 O(n)。
    空间复杂度：
        这里的空间复杂度和递归使用的栈空间有关，这里递归层数不超过 n，故空间复杂度为 O(n)。

【迭代】
    实际上递归就是调用了栈 => 所以基本上递归都能转化成栈 
    

*/