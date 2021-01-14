// 题目及解题思路更新到‘解题思路.md’

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
function lowestCommonAncestor(root: TreeNode | null, p: TreeNode | null, q: TreeNode | null): TreeNode | null {

	if(root){
        const mVal = root.val;
        const pVal = p ? p.val : null;
        const qVal = q ? q.val : null;
        if(mVal > pVal && mVal > qVal) return lowestCommonAncestor(root.left, p, q);
        if(mVal < pVal && mVal < qVal) return lowestCommonAncestor(root.right, p, q)
    }

    return root;
};