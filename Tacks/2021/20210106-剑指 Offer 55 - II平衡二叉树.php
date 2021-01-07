<?php

/*
 * @Descripttion: 周三，今天北京的风超级大，真特么爽 冻死个人
 * @Author: tacks321@qq.com
 * @Date: 2021-01-06 09:15:41
 * @LastEditTime: 2021-01-06 18:36:13
 */


/*
剑指 Offer 55 - II. 平衡二叉树


输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。

示例 1:

给定二叉树 [3,9,20,null,null,15,7]

    3
   / \
  9  20
    /  \
   15   7
返回 true 。

示例 2:

给定二叉树 [1,2,2,3,3,null,null,4,4]

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
返回 false 。

限制：

1 <= 树的结点个数 <= 10000
注意：本题与主站 110 题相同：https://leetcode-cn.com/problems/balanced-binary-tree/

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ping-heng-er-cha-shu-lcof
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
// 【先序遍历 自上向下】
class Solution1 {

    /**
     * @param TreeNode $root
     * @return Boolean
     */
    // 执行用时：20 ms, 在所有 PHP 提交中击败了36.36%的用户
    // 内存消耗：17.4 MB, 在所有 PHP 提交中击败了87.88%的用户
    // 先序遍历 自顶向下
    function isBalanced($root) {
        if($root == null) {
            return true;
        } else {
            // 对当前的节点 前序遍历
            // 如果当前节点的左右子树的高度差不超过1，再去分别遍历左子树跟右子树
            return abs($this->height($root->left) - $this->height($root->right)) <= 1 
                && $this->isBalanced($root->left)
                && $this->isBalanced($root->right);
        }
    }

    // 计算树的高度
    function height($p) {
        if($p == null) {
            return 0;
        } else {
            // 非空节点
            return max($this->height($p->left), $this->height($p->right)) + 1;
        }
    }

}


// 【后序遍历 自下而上】
class Solution2 {

    /**
     * @param TreeNode $root
     * @return Boolean
     */
    // 执行用时：16 ms, 在所有 PHP 提交中击败了48.48%的用户
    // 内存消耗：17.7 MB, 在所有 PHP 提交中击败了21.21%的用户
    function isBalanced($root) {
        return $this->recur($root) !== -1;
    }

    function recur($node) {
        if($node == null) {
            return 0;
        }
        
        $leftHeight = $this->recur($node->left);
        if($leftHeight == -1) return -1;
        $rightHeight = $this->recur($node->right);
        if($rightHeight == -1) return -1;

        // 高度差
        if(abs($leftHeight - $rightHeight) > 1 ) {
            return -1;
        } else {
            // 返回当前节点的高度
            return max($leftHeight, $rightHeight) + 1;
        }
    }
     

}



/*

这个是原题 没跑了。

[110] 平衡二叉树 https://leetcode-cn.com/problems/balanced-binary-tree/


平衡二叉树 概念：二叉树的每个节点的左右子树的高度差绝对值不超过1


【1】先序遍历 自上向下
    获取左子树高度，获取右子树高度，比较两个子树高度差，是否小于等于1
    先序遍历 当前子树的左子树是否平衡
    先序遍历 当前子树的右子树是否平衡
    如果所有子树都平衡，整棵树才平衡

    但是从上而下，每次都要计算树的高度，就会重复计算，时间复杂度较高。

    时间复杂度 O(nlogn) 。 二叉树的高度计算复杂度O(logN),每层遍历需要N个，大概。。。。
    空间复杂度 O(n)     。 需要用到递归栈
    
   
【2】后续遍历 + 剪枝 自下而上 

    剪枝 这个词听到次数不多，感觉牛逼NB 🐂
    
    思路：自下而上返回子树的深度，如果判定某个子树不是平衡树则进行 剪枝，直接向上返回。

    时间复杂度 O(n) N为树的节点数；最差情况下，需要递归遍历树的所有节点。
    空间复杂度 O(n) 需要用到递归栈





*/