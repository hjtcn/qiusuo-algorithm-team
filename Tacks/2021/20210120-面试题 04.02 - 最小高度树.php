<?php
/*
 * @Descripttion: 周三一过，周五就不远了 啊啊啊
 * @Author: tacks321@qq.com
 * @Date: 2021-01-20 18:22:37
 * @LastEditTime: 2021-01-20 18:58:02
 */


/*

面试题 04.02. 最小高度树

给定一个有序整数数组，元素各不相同且按升序排列，编写一个算法，创建一棵高度最小的二叉搜索树。

示例:
给定有序数组: [-10,-3,0,5,9],

一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：

          0 
         / \ 
       -3   9 
       /   / 
     -10  5 

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-height-tree-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
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
class Solution {

    /**
     * @param Integer[] $nums
     * @return TreeNode
     */
    function sortedArrayToBST($nums) {
        return $this->buildsubtree($nums, 0 , count($nums));
    }   

    // 递归构建
    function buildsubtree($nums, $left, $right) {
        // 如果都是空的就结束
        if( $left == $right) {
            return null;
        }
        // 以中间节点为根节点，一分为二
        // 根据mid划分左子数组和右子数组。
        $mid = floor(($left+$right)/2);

        $node= new TreeNode($nums[$mid]);
        // 左子数组构建左子树
        $node->left = $this->buildsubtree($nums, $left, $mid);
        // 右子数组构建右子树
        $node->right= $this->buildsubtree($nums, $mid+1, $right);

        return $node;
    }
}


/*

[题目]
    给定一个有序整数数组，元素各不相同且按升序排列，编写一个算法，创建一棵高度最小的二叉搜索树。

[思考]
    问题：树的高度是指?    => 树从根结点开始往下数，叶子结点所在的最大层数
    问题：二叉搜索树是?    => 可以是空树；没有节点的值是相等的；如果任意节点存在 左子树 根节点 右子树 那么 左 < 根 < 右

画画图就很容易理解了
    [1,2,3,4,5]
        那么3作为根节点
         3
    1 2     4 5   然后再一分二
    

时间复杂度O(n)
*/