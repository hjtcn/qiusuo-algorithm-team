// 112. 路径总和
// 给定一个二叉树和一个目标和，判断该树中是否存在根节点到叶子节点的路径，这条路径上所有节点值相加等于目标和。

// 说明: 叶子节点是指没有子节点的节点。

// 示例: 
// 给定如下二叉树，以及目标和 sum = 22，

//               5
//              / \
//             4   8
//            /   / \
//           11  13  4
//          /  \      \
//         7    2      1
// 返回 true, 因为存在目标和为 22 的根节点到叶子节点的路径 5->4->11->2。


// ============================================================
// ===                                                      ===
// ===      状态：通过,执行用时: 100 ms,内存消耗: 43.5 MB     ===
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

function hasPathSum(root: TreeNode | null, sum: number): boolean {
    let hasPathSum = false;

    function mapTree (node: TreeNode | null, currentTotal: number) {

        if(hasPathSum) return;

        if(node){
            const num = currentTotal + node.val;
            if(!node.left && !node.right){
                hasPathSum = num === sum;
            }else{
                node.left && mapTree(node.left, num);
                node.right && mapTree(node.right, num);
            }
        }

    }

    mapTree(root, 0);
    return hasPathSum;
};