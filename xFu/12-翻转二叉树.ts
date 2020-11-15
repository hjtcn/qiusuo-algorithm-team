
// 226. 翻转二叉树
// 翻转一棵二叉树。

// 示例：

// 输入：

//      4
//    /   \
//   2     7
//  / \   / \
// 1   3 6   9
// 输出：

//      4
//    /   \
//   7     2
//  / \   / \
// 9   6 3   1
// 备注:
// 这个问题是受到 Max Howell 的 原问题 启发的 ：

// 谷歌：我们90％的工程师使用您编写的软件(Homebrew)，但是您却无法在面试时在白板上写出翻转二叉树这道题，这太糟糕了。


// ============================================================
// ===                     自己想的代码                      ===
// ===       状态：通过,执行用时: 88 ms,内存消耗: 40.3 MB     ===
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

function invertTree(root: TreeNode | null): TreeNode | null {

    function invertNode (node: TreeNode | null){

        if(node === null) return;

        // 修改当前结点的左右子树位置
        const tempNode = node.left;
        node.left = node.right;
        node.right = tempNode;

        node.left && invertNode(node.left);
        node.right && invertNode(node.right);
    }

    invertNode(root);

    return root;
};

// ============================================================
// ===                   参考他人代码后想的                   ===
// ===       状态：通过,执行用时: 84 ms,内存消耗: 40.4 MB     ===
// ============================================================

function invertTree_1(root: TreeNode | null): TreeNode | null {

    // 只有当前结点不为null才执行之后的翻转，否则直接返回root(此时root为null)
    if(root !== null){
        // 先判断左子树，再判断右子树。最后将结果调换位置，实现翻转
        const left =  invertTree(root.right);
        const right = invertTree(root.left);

        root.left = left;
        root.right = right;
    }

    return root;
};