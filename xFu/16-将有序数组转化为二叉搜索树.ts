// 108. 将有序数组转换为二叉搜索树
// 将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。

// 本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

// 示例:

// 给定有序数组: [-10,-3,0,5,9],

// 一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：

//       0
//      / \
//    -3   9
//    /   /
//  -10  5


// ============================================================
// ===                                                      ===
// ===       状态：通过,执行用时: 96 ms,内存消耗: 42.2 MB     ===
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

function sortedArrayToBST(nums: number[]): TreeNode | null {
    let root = null;

    function mapArray(start:number, end:number): TreeNode | null{
        const midIndex = Math.floor((start + end)/2);

        return start < end 
        ? {
            val: nums[midIndex],
            left: mapArray(start, midIndex),
            right: mapArray(midIndex + 1, end),
        }
        : null
    }

    root = mapArray(0, nums.length);

    return root;
};