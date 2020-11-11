// 104. 二叉树的最大深度
// 给定一个二叉树，找出其最大深度。

// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

// 说明: 叶子节点是指没有子节点的节点。

// 示例：
// 给定二叉树 [3,9,20,null,null,15,7]，

//     3
//    / \
//   9  20
//     /  \
//    15   7
// 返回它的最大深度 3 。


// ============================================================
// ===                     自己想的代码                      ===
// ===       状态：通过,执行用时: 92 ms,内存消耗: 42.7 MB     ===
// ============================================================

/**
 * Definition for a binary tree node.
 * class TreeNode {
 *     val: number
 *     left: TreeNode | null
 *     right: TreeNode | null
 *     constructor(val?: number, left?: TreeNode | null, right?: TreeNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.left = (left===undefined ? null : left)
 *         this.right = (right===undefined ? null : right)
 *     }
 * }
 */

interface TreeNode {
    val: number
    left: TreeNode | null
    right: TreeNode | null
 }

function maxDepth(root: TreeNode | null): number {
    let maxDepth = 0;

    function judeDepth (currentDepth: number, node: TreeNode | null) {
        if(node !== null){
            currentDepth ++
            judeDepth(currentDepth, node.left);
            judeDepth(currentDepth, node.right);
        }else{
            maxDepth = Math.max(maxDepth, currentDepth);
        }
    }

    judeDepth(0, root);

    return maxDepth;
};