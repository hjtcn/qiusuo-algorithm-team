<?php

/*
 * @Description: 今天冲冲冲，把这周目标完成一下
 * @Author: Tacks
 * @Date: 2020-11-19 15:17:17
 */
/*
 * @lc app=leetcode.cn id=面试题 04.03. lang=php
 *
 * [面试题 04.03.]  特定深度节点链表
 *
 * https://leetcode-cn.com/problems/list-of-depth-lcci/
 *
 *
 * 给定一棵二叉树，设计一个算法，创建含有某一深度上所有节点的链表（比如，若一棵树的深度为 D，则会创建出 D 个链表）。返回一个包含所有深度的链表的数组。
 * 
 * 
 * 
 * 
 * 示例
 * 
    输入：[1,2,3,4,5,null,7,8]

            1
        /  \ 
        2    3
        / \    \ 
        4   5    7
    /
    8

    输出：[[1],[2,3],[4,5,7],[8]]
 
*/

/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     public $val = null;
 *     public $left = null;
 *     public $right = null;
 *     function __construct($value) { $this->val = $value; }
 * }
 */
/**
 * Definition for a singly-linked list.
 * class ListNode {
 *     public $val = 0;
 *     public $next = null;
 *     function __construct($val) { $this->val = $val; }
 * }
 */
class Solution {

    public $list = [];
    /**
     * @param TreeNode $tree
     * @return ListNode[]
     */
    // 执行用时：4 ms, 在所有 PHP 提交中击败了84.62%的用户
    // 内存消耗：15.3 MB, 在所有 PHP 提交中击败了7.69%的用户
    function listOfDepth($tree) {
        if(!$tree) {
            return [];
        }
        $this->dfs($tree, 0);
        return $this->list;
    }

    function dfs($node, $level) {
        if(!$node) {
            return null;
        }
        // 处理某层的时候
        if(!isset($this->list[$level])) {
            $this->list[$level] = new ListNode($node->val);
        }else{
            // 如果已经加入，那么就头插入节点链表
            $temp = $this->list[$level];
            $this->list[$level] = new ListNode($node->val);
            $this->list[$level]->next = $temp;
        }
        $this->dfs($node->right, $level+1);        
        $this->dfs($node->left, $level+1);        
    }
}

/*

【深度优先遍历】DFS
    递归处理

【广度优先遍历】BFS


*/