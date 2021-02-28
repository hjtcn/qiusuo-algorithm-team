<?php

// 2021 第一题

/*

【剑指 Offer 34. 二叉树中和为某一值的路径 】

输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。

 

示例:
给定如下二叉树，以及目标和 sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:

[
   [5,4,11,2],
   [5,8,4,5]
]
 

提示：

节点总数 <= 10000
注意：本题与主站 113 题相同：https://leetcode-cn.com/problems/path-sum-ii/

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof
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

    private $result = []; // 结果路径
    /**
     * @param TreeNode $root
     * @param Integer $sum
     * @return Integer[][]
     */
    // 递归处理
    // 执行用时：24 ms, 在所有 PHP 提交中击败了17.86%的用户
    function pathSum($root, $sum) {
        $this->dfs($root, $sum, []);
        return $this->result;
    }

    function dfs($node, $sum, $path) {
        if($node == null) {
            return [];
        }
        $path[] = $node->val; // 记录当前路径
        // 递归终止条件： 遍历到叶子节点
        if($node->left == null && $node->right == null) {
            // 满足条件
            if($node->val == $sum) {
                $this->result[] = $path;
                return ;
            }
        }
        // 继续向下递归
        $this->dfs($node->left, $sum - $node->val, $path);
        $this->dfs($node->right, $sum - $node->val, $path);
    }
}


class Solution2 {

    /**
     * @param TreeNode $root
     * @param Integer $sum
     * @return Integer[][]
     */
    // DFS
    // 执行用时：20 ms, 在所有 PHP 提交中击败了35.71%的用户
    function pathSum($root, $sum) {
        // 边界判断
        if($root == null) {
            return [];
        }
        // 节点 路径 当前值
        $queue[] = [$root,[$root->val], $root->val];
        $result = [];

        while(!empty($queue)) {
            // 批量赋值
            list($node, $path, $cur_sum) = array_shift($queue);

            // 符合要求
            if($node->left == null && $node->right == null && $cur_sum == $sum) {
                $result[] = $path;
            }

            // 不等于空
            if($node->left != null) {
                $path[] = $node->left->val;
                // 更新队列
                $queue[] = [$node->left, $path, $cur_sum + $node->left->val];
                array_pop($path); // 移除左边这个路径
            }

            if($node->right != null) {
                $path[] = $node->right->val;
                $queue[] = [$node->right, $path, $cur_sum + $node->right->val];
            }
        }
        return $result;
    }
}