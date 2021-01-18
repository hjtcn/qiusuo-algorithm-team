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
    let notBalanced = false;

    function mapTree(node: TreeNode | null): number {
        if(notBalanced) return;
        if(node) {
            const left = mapTree(node.left);
            if(notBalanced) return;

            const right = mapTree(node.right);
            if(notBalanced) return;

            if(Math.abs(+left - +right) > 1){
                notBalanced = true;
            }

            return Math.max(+left, +right) + 1;
        }

        return 0;
    }

    mapTree(root);

    return !notBalanced;
};