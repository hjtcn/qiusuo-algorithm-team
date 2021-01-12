// 剑指 Offer 55 - II. 平衡二叉树
// 输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。



// 示例 1:

// 给定二叉树 [3,9,20,null,null,15,7]

//     3
//    / \
//   9  20
//     /  \
//    15   7
// 返回 true 。

// 示例 2:

// 给定二叉树 [1,2,2,3,3,null,null,4,4]

//        1
//       / \
//      2   2
//     / \
//    3   3
//   / \
//  4   4
// 返回 false 。



// 限制：

// 1 <= 树的结点个数 <= 10000


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

function isBalanced(root: TreeNode | null): boolean {
    let status = true;

    function DFS(node: TreeNode | null, level: number) {
        // 如果状态已经是不平衡了就没必要继续遍历了
        if(!status) return 0;

        if(node === null){
            return level;
        }else {
            level++;
            const leftLevel = DFS(node.left, level);
            const rightLevel = DFS(node.right, level);
            const balanceNum = leftLevel - rightLevel;

            if(Math.abs(balanceNum) > 1){
                status = false;
            }

            return balanceNum > 0 ? leftLevel : rightLevel;
        }
    }

    DFS(root, 1);
    return status;
}; 