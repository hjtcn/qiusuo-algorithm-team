<?php
/*
 * @Descripttion: 周二搞起
 * @Author: tacks321@qq.com
 * @Date: 2021-06-15 17:34:54
 * @LastEditTime: 2021-06-15 18:51:11
 */


/*
 * @lc app=leetcode.cn id=111 lang=php
 *
 * [111] 二叉树的最小深度
 *
 * https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/description/
 *
 * algorithms
 * Easy (47.85%)
 * Likes:    524
 * Dislikes: 0
 * Total Accepted:    228.3K
 * Total Submissions: 476.7K
 * Testcase Example:  '[3,9,20,null,null,15,7]'
 *
 * 给定一个二叉树，找出其最小深度。
 * 
 * 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
 * 
 * 说明：叶子节点是指没有子节点的节点。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：root = [3,9,20,null,null,15,7]
 * 输出：2
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：root = [2,null,3,null,4,null,5,null,6]
 * 输出：5
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 树中节点数的范围在 [0, 10^5] 内
 * -1000 
 * 
 * 
 */

// @lc code=start
//  Definition for a binary tree node.
class TreeNode {
    public $val = null;
    public $left = null;
    public $right = null;
    function __construct($val = 0, $left = null, $right = null) {
        $this->val = $val;
        $this->left = $left;
        $this->right = $right;
    }
}



class Solution {

    /**
     * @param TreeNode $root
     * @return Integer
     */
    // 深搜DFS
    // 执行用时：452 ms, 在所有 PHP 提交中击败了11.90%的用户
    // 内存消耗：46.8 MB, 在所有 PHP 提交中击败了14.28%的用户
    function minDepth($root) {
        // 递归终止条件
        if($root == null) {
            return 0;
        }
        
        // 左右子树都存在
        if($root->left && $root->right) {
            return 1 + min($this->minDepth($root->left), $this->minDepth($root->right));
        } elseif($root->left) {
            // 如果只有左子树
            return 1 + $this->minDepth($root->left);
        } elseif($root->right) {
            // 如果只有右子树
            return 1 + $this->minDepth($root->right);
        } else {
            // 如果是叶子姐弟啊
            return 1;
        }
    }
}


class Solution2 {

    /**
     * @param TreeNode $root
     * @return Integer
     */
    // 广搜BFS
    // 执行用时：372 ms, 在所有 PHP 提交中击败了66.67%的用户
    // 内存消耗：46.6 MB, 在所有 PHP 提交中击败了78.57%的用户
    function minDepth($root) {
        // 边界情况
        if($root == null) {
            return 0;
        }
        if($root->left == null && $root->right == null) {
            return 1;
        }

        $queue = [$root]; // 队列 根节点入队
        $depth = 0; // 最小深度
        while($queue) {
            $depth++; // 每一层都++
            $size = count($queue);
            // 层次遍历
            for ($i=0; $i < $size; $i++) { 
                $node = array_shift($queue);  // 队列 先进先出 移除头部节点
                // 遇到第一个叶子节点，直接返回高度 也就是所在层数
                if($node->left == null && $node->right == null) {
                    return $depth;
                }
                if($node->left != null ) {
                    $queue[] = $node->left;   // 入队
                }
                if($node->right != null) {
                    $queue[] = $node->right;  // 入队
                }
            }
        }
        return $depth;
    }
}
// @lc code=end



/*
【题目】
    二叉树的最小深度
        这个题出现过多次了    在二叉树训练的时候，这种题就是最基本的类型
        目前训练的是深搜广搜  二叉树的类型题目仍然是不错的典型题
    
    知识点
        二叉树结构
        二叉树的最小深度（从根节点到最近叶子节点的最短路径上的节点数量）
        二叉数的最大深度（同类型题目） [[104]二叉树的最大深度](https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/)
        二叉树的遍历（前序、层次）
        深搜 广搜
        递归
        队列
【解析】

    最小深度：需要增加代码，判断是叶子节点。
    最大深度：能够保证最后一个节点绝对是叶子节点
    
    1、DFS 深搜
        前序遍历整个树，记录最小深度
            对于每一个非叶子节点，分别计算左右子树的最小叶子节点深度，递归处理

        复杂度
            时间复杂度：O(N)，其中 N 是树的节点数。对每个节点访问一次。
            空间复杂度：O(H)，其中 H 是树的高度。空间复杂度主要取决于递归时栈空间的开销
                                最坏情况下，树呈现链状，空间复杂度为 O(N)。平均情况下树的高度与节点数的对数正相关，空间复杂度为 O(logN)。

 
    2、BFS 广搜
        层次遍历整个树，当我们找到一个叶子节点，直接返回深度。
            主要就是队列的使用，将每一层的节点从左到右先入队，然后依次出队，获取深度。      
            
        复杂度
            时间复杂度：O(N)，其中 NN 是树的节点数。对每个节点访问一次。
            空间复杂度：O(N)，其中 N 是树的节点数。空间复杂度主要取决于队列的开销，队列中的元素个数不会超过树的节点数。

 
【参考】
    https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/solution/tu-jie-dfs-xie-liao-si-ban-bfs-xie-liao-yi-ban-by-/
    https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/solution/111-er-cha-shu-de-zui-xiao-shen-du-di-gu-ztum/


*/