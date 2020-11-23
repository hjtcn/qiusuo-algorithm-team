// 100. 相同的树
// 给定两个二叉树，编写一个函数来检验它们是否相同。

// 如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

// 示例 1:

// 输入:       1         1
//           / \       / \
//          2   3     2   3

//         [1,2,3],   [1,2,3]

// 输出: true
// 示例 2:

// 输入:      1          1
//           /           \
//          2             2

//         [1,2],     [1,null,2]

// 输出: false
// 示例 3:

// 输入:       1         1
//           / \       / \
//          2   1     1   2

//         [1,2,1],   [1,1,2]

// 输出: false




// ============================================================
// ===                     自己想的代码                      ===
// ===       状态：通过,执行用时: 76 ms,内存消耗: 40.4 MB     ===
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

function isSameTree(p: TreeNode | null, q: TreeNode | null): boolean {
    let stack: (number | null)[] = [];

    // 递归取出树的所有值
    function mapTree (node: TreeNode | null) {
        if(node !== null){
            stack.push(node.val);
            mapTree(node.left);
            mapTree(node.right);
        }else{
            stack.push(null);
        }
    }

    // 遍历取出p树的所有值
    mapTree(p);
    // 遍历取出q树的所有值
    mapTree(q);

    // 若列表的长度不是偶数，就表示一定不相同
    if(stack.length % 2 !== 0){
        return false;
    }
    
    // 取列表的中间数坐标
    let midIndex = stack.length / 2;

    // 判断列表数是否相同
    for(let i = 0;i<midIndex;i++){
        if(stack[i] !== stack[i+midIndex]){
            return false;
        }
    }

    return true;
};