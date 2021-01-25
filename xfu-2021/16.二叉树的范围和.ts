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

function rangeSumBST(root: TreeNode | null, low: number, high: number): number {
    if(root) {
        const value = root.val;

        // 若中间结点的值比最小值小就递归判断它的右子结点（因为右子结点一定比中间结点大）
        if(value < low) return rangeSumBST(root.right, low, high);

        // 若中间结点的值比最大值大就递归判断它的左子结点（因为左子结点一定比中间结点小）
        if(value > high) return rangeSumBST(root.left, low, high);

        return value + rangeSumBST(root.left, low, high) + rangeSumBST(root.right, low, high);
    }
    return 0;
};