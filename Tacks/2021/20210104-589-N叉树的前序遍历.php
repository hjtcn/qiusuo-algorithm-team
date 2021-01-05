<?php
/*
 * @Descripttion: 2021第一个工作日 冲冲冲啊 
 * @Author: tacks321@qq.com
 * @Date: 2021-01-04 16:27:21
 * @LastEditTime: 2021-01-05 13:53:35
 */

/*
 * @lc app=leetcode.cn id=589 lang=php
 *
 * [589] N叉树的前序遍历
 *
 * https://leetcode-cn.com/problems/n-ary-tree-preorder-traversal/description/
 *
 * algorithms
 * Easy (74.06%)
 * Likes:    128
 * Dislikes: 0
 * Total Accepted:    60.1K
 * Total Submissions: 81.2K
 * Testcase Example:  '[1,null,3,2,4,null,5,6]'
 *
 * 给定一个 N 叉树，返回其节点值的前序遍历。
 * 
 * 例如，给定一个 3叉树 :
 * 
 * 返回其前序遍历: [1,3,5,6,2,4]。
 * 
 * 
 * 
 * 说明: 递归法很简单，你可以使用迭代法完成此题吗?
 */

// @lc code=start
/**
 * Definition for a Node.
 * class Node {
 *     public $val = null;
 *     public $children = null;
 *     function __construct($val = 0) {
 *         $this->val = $val;
 *         $this->children = array();
 *     }
 * }
 */

class Solution1 {
    /**
     * @param Node $root
     * @return integer[]
     */
    // 执行用时：20 ms, 在所有 PHP 提交中击败了46.43%的用户
    // 内存消耗：19 MB, 在所有 PHP 提交中击败了41.07%的用户
    // 递归思想
    function preorder($root) {
        $result = [];
        // if($root == null) {
        //     return $result;
        // }
        $this->_helper($root, $result);

        return $result;
    }

    // 辅助函数
    private function _helper($node, &$result)
    {
        if($node == null) {
            return $result;
        }
        $result[] = $node->val; // 中序遍历
        // 递归处理
        /*
        // 如果是二叉树，则递归代码如下
            $this->_helper($node->left, $result);
            $this->_helper($node->right, $result);
        */
        // 循环遍历子树
        foreach($node->children as $child) {
            $this->_helper($child, $result);
        }
    }
}



class Solution2 {
    /**
     * @param Node $root
     * @return integer[]
     */
    // 迭代思想
    // 执行用时：12 ms, 在所有 PHP 提交中击败了98.21%的用户
    // 内存消耗：19 MB, 在所有 PHP 提交中击败了41.07%的用户
    function preorder($root) {
        $result = [];
        if($root === null) {
            return $result;
        }
        // 采用栈
        $stack[] = $root;
        while($stack) {
            $node = array_pop($stack);
            $result[] = $node->val; // 根节点

            // 二叉树 处理方式
            /*
            if($node->right) { $stack[] = $node->right; }
            if($node->left)  { $stack[] = $node->left; }
            */
            // N叉树处理方式
            if($node->children !== null) {
                // 这里的子节点，应该从右开始压入栈。
                $nodeList = array_reverse($node->children);
                foreach($nodeList as $child) {
                    $stack[] = $child; // 压入栈中
                }
            }
        }
        return $result;
    }
}


// @lc code=end



/*

前序遍历

【递归】【迭代】


【思路】
二叉树 => N叉树

1、我们知道二叉树前序遍历如何编写，先根节点-左子树-右子树
2、现在由2=>N的变化，体现在子树中，原本使用两个指针，分别指向左子树和右子树，现在采用children表示子树数组

那么可以套用二叉树的前序遍历代码模版，在关键处进行调整。





*/