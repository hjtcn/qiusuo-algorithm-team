// 404. 左叶子之和
// 计算给定二叉树的所有左叶子之和。

// 示例：

//     3
//    / \
//   9  20
//     /  \
//    15   7

// 在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24

// ============================================================
// ===                                                      ===
// ===       状态：通过,执行用时: 80 ms,内存消耗: 40.6 MB     ===
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

function sumOfLeftLeaves(root: TreeNode | null): number {
    let leftCount = 0;

    function mapTree (node: TreeNode | null, position: 'left' | 'right'){
        if(node){
            if(!node.left && !node.right && position === 'left'){
                leftCount += node.val;
            }else{
                mapTree(node.left, 'left');
                mapTree(node.right, 'right');
            }
        }
        
    }

    mapTree(root, 'right');

    return leftCount;
};