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

function sortedArrayToBST(nums: number[]): TreeNode | null {
    if(nums.length){
        const midNum = Math.floor(nums.length / 2);
        const node = new TreeNode();

        node.val = nums[midNum];
        node.left = sortedArrayToBST(nums.slice(0, midNum));
        node.right = sortedArrayToBST(nums.slice(midNum+1));

        return node;
    }
    return null;
};