// 105. 从前序与中序遍历序列构造二叉树
// 根据一棵树的前序遍历与中序遍历构造二叉树。

// 注意:
// 你可以假设树中没有重复的元素。

// 例如，给出

// 前序遍历 preorder = [3,9,20,15,7]
// 中序遍历 inorder = [9,3,15,20,7]
// 返回如下的二叉树：

//     3
//    / \
//   9  20
//     /  \
//    15   7

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

function buildTree(preorder: number[], inorder: number[]): TreeNode | null {
    if(!preorder.length) return null;

    // 获取到前序的第一个数（根）在中序中的位置
    const mid = inorder.indexOf(preorder[0]);

    return new TreeNode(
        // 保存根的值
        preorder[0],
        // 构建左子树。
        // 前序数组从第二个数开始取，取到和中序数组长度一致为止。
        // 中序数组取到前序第一个数所在位置为止
        buildTree(preorder.slice(1, mid+1), inorder.slice(0, mid)),
        // 构建右子树
        buildTree(preorder.slice(mid+1), inorder.slice(mid+1))
        )
    
};
