<?php
/*
 * @Descripttion: 复习周-做原题  周一快乐  口腔溃疡疼死个人哇哇哇哇！😭
 * @Author: tacks321@qq.com
 * @Date: 2021-01-18 18:03:03
 */


/*
 * @lc app=leetcode.cn id=110 lang=php
 *
 * [110] 平衡二叉树
 *
 * https://leetcode-cn.com/problems/balanced-binary-tree/description/
 *
 * algorithms
 * Easy (55.09%)
 * Likes:    568
 * Dislikes: 0
 * Total Accepted:    165.8K
 * Total Submissions: 300.8K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，判断它是否是高度平衡的二叉树。
 * 
 * 本题中，一棵高度平衡二叉树定义为：
 * 
 * 
 * 一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。
 * 
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：root = [3,9,20,null,null,15,7]
 * 输出：true
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：root = [1,2,2,3,3,null,null,4,4]
 * 输出：false
 * 
 * 
 * 示例 3：
 * 
 * 
 * 输入：root = []
 * 输出：true
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 树中的节点数在范围 [0, 5000] 内
 * -10^4 
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
class Solution1 {

    /**
     * @param TreeNode $root
     * @return Boolean
     */
    // 执行用时：24 ms, 在所有 PHP 提交中击败了23.33%的用户
    // 内存消耗：17.3 MB, 在所有 PHP 提交中击败了75.00%的用户
    // 自顶向下 
    function isBalanced($root) {
        if($root == null) {
            return true;
        }else{
            // 当前节点前序遍历
            // 如果当前节点的左右子树高度差不超过1，那么再分别去遍历左子树和右子树。
            return abs($this->height($root->left) - $this->height($root->right)) <= 1  
                    && $this->isBalanced($root->left)
                    && $this->isBalanced($root->right);
        }
    }

    // 计算节点的高度
    function height($node) {
        if($node == null) {
            return 0;
        }else{
            // 非空节点进行处理
            return max($this->height($node->left), $this->height($node->right)) + 1;
        }
    }

}

class Solution2 {

    /**
     * @param TreeNode $root
     * @return Boolean
     */
    // 执行用时：8 ms, 在所有 PHP 提交中击败了100.00%的用户
    // 内存消耗：17.6 MB, 在所有 PHP 提交中击败了48.33%的用户
    // 自底向上
    function isBalanced($root) {
        return $this->height($root) >=0 ;
    }

    function height($node) {
        // 当前节点为空，直接返回
        if($node == null) {
            return 0;
        }
        // 计算左右子树的高度
        $leftHeight = $this->height($node->left); // 向左递归遍历，计算高度
        $rightHeight= $this->height($node->right);// 向右递归遍历，计算高度
        // 高度差超过1 或者节点不存在
        if($leftHeight == -1 || $rightHeight == -1 || abs($leftHeight - $rightHeight) > 1) {
            return -1;
        }else{
            // 返回当前节点高度。
            return max($leftHeight, $rightHeight) + 1;
        }
    }

}
// @lc code=end



/*

题目：
    给定一个二叉树，判断它是否是高度平衡的二叉树。
 
思考：
    什么是二叉树？   每个节点最多只有两个分支，也就是左节点、右节点
    什么是高度平衡？ 一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 
    有什么特性？     一个平衡二叉树的左右子树也分别是平衡二叉树。（有点套娃的感觉🙉，看到这种递归的思路就出来了）
    如何去递归？     自顶向下or自底向上

[解法1]自顶向下
    比如二叉树中任意一个节点，我们如何计算节点node的高度
        当node==null 也就是空节点 
            height(node) = 0
        当node!=null 非空节点
            height(node) = max(height(node.left), height(node.right)) + 1 
    那么当我们知道一个节点如何去计算高度，便可以用特点，高度平衡，左右节点差的绝对值<=1
        我们可以前序遍历（前序遍历也是一个递归的过程）
            先遍历根节点，判断根节点是否平衡（每一次平衡都要去计算当前节点的左右子树高度差，也就是一个递归的过程）
            然后遍历左节点，判断左子树是否平衡
            最后遍历右节点，判断右子树是否平衡
    我们是从根节点向下，也即是叶子节点，不断递归遍历
所以这个方法的递归是自顶向下！

时间复杂度：O(n^2) n为二叉树的节点个数 （因为它前序遍历如果是满二叉树，最多要遍历O(n)。另外计算高度，平均每次都要调用log(n)次，最差链式要调用n次）
空间复杂度：O(n)， n是二叉树中的节点个数。空间复杂度主要取决于递归调用的层数，递归调用的层数不会超过 n。

[解法2]自底向上
    由于，自顶向下，在计算同一个节点的高度会重复计算，导致时间复杂度偏高，这里采用自底向上，就会好一些，对每个节点只需要height计算一次

    自底向上的做法，类似后序遍历
        遍历当前节点，首先递归判断左右子树是否平衡，再判断当前根节点是否平衡
        如果一颗子树是平衡的就返回高度，否则返回-1。
    如果一个子树不是平衡的，那么整个二叉树也不是平衡的。

时间复杂度:O(n) 每个节点的计算高度和判断是否平衡都只需要处理一次，最坏情况下需要遍历二叉树中的所有节点，因此时间复杂度是 O(n)
空间复杂度:O(n)
###########################
看到这里我突然意识到，自顶向下，和自底向上的区别
    自顶向下，前序遍历。
        就是你要从上向下依次判断，当前节点子树的高度差，然后一直到叶子节点。
    自底向上，后序遍历。
        就是你判断子树，直接去叶子节点找，如果有一个子树不平衡，直接结束。时间上会节约掉高度重复计算的过程。
###########################





👓 看到有一个很好的评论

    自底向上的递归就是及时止损 
    自顶向下的递归就是不管三七二十一先把所有的高度算一遍

所以也能看出来自底向上的递归性能是比较好的


 */


 // 重新调整了一下
class Solution3 {

    /**
     * @param TreeNode $root
     * @return Boolean
     */
    // 执行用时：8 ms, 在所有 PHP 提交中击败了100.00%的用户
    // 内存消耗：17.6 MB, 在所有 PHP 提交中击败了48.33%的用户
    // 自底向上
    function isBalanced($root) {
        return $this->height($root) >=0 ;
    }

    function height($node) {
        // 当前节点为空，直接返回
        if($node == null) {
            return 0;
        }
        ////////////////////////////////【递归进行剪枝：左边递归-1，直接结束】////////////////////////////////////////
        // 计算左右子树的高度
        $leftHeight = $this->height($node->left); // 向左递归遍历，计算高度
        // 剪枝： 提前结束递归。  如果左子树已经返回-1 就不需要再递归右子树了，直接返回-1
        if($leftHeight == -1) {
            return -1;
        }
        $rightHeight= $this->height($node->right);// 向右递归遍历，计算高度
        if($rightHeight == -1) {
            return -1;
        }
        ///////////////////////////////////////////////////////////////////////////////////


        // 高度差超过1 
        if(abs($leftHeight - $rightHeight) > 1) {
            return -1;
        }else{
            // 返回当前节点高度。
            return max($leftHeight, $rightHeight) + 1;
        }
    }

}


