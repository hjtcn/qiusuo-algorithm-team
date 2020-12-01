// 145. 二叉树的后序遍历
// 给定一个二叉树，返回它的 后序 遍历。

// 示例:

// 输入: [1,null,2,3]  
//    1
//     \
//      2
//     /
//    3 

// 输出: [3,2,1]
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？


/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number[]}
 */
var postorderTraversal = function(root) {
    let res=[]
    let helper=(root)=>{
        if(!root) return null;
        helper(root.left)
        helper(root.right)
        res.push(root.val)
    }
    helper(root)
    return res
};
//递归只需要知道左-右-根就行
/*
    时间复杂度：O(N)
    取决于树的节点数
    空间复杂度：O(N)
    包含树的节点数数

*/



var postorderTraversal = function(root) {
    let res=[]
    let stack=[root]
    while(stack.length){   
        let root=stack.pop()
        root&&res.unshift(root.val)
        root&&stack.push(root.left)
        root&&stack.push(root.right)
    }
    return res
};
//迭代这个思路其实有点酷，我自己做的时候净想着怎么左->右->根了
//看了题解才发现将根右左的val进行unshift就是左右根了
// 而根右左如何获取呢？
//根直接通过root.val获取就行
//右->左的实现方案：先入栈左子树，再入栈右子树，出栈就变成先右后左了

/*
    时间复杂度：O(N)
    取决于树的节点数
    空间复杂度：O(N)
    包含树的节点数数

*/