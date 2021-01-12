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



/*
 * @lc app=leetcode.cn id=110 lang=javascript
 *
 * [110] 平衡二叉树
 */

// @lc code=start
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
 * @return {boolean}
 */
var isBalanced = function (root) {
    if (!root) return true
    let getHeight = (node) => {
        if (!node) return -1;
        return 1 + Math.max(getHeight(node.left), getHeight(node.right))
    }
    return Math.abs(getHeight(root.left) - getHeight(root.right)) <= 1 && isBalanced(root.left) && isBalanced(root.right)
};
// @lc code=end
/** 
 是原题，做题的时候，完全没有原题的思路了，绝望，没有记忆点。
 递归获取高度。满足当前root的左右子树高度差<=1且递归调用左右子树
   时间复杂度：O(nlogn)
    getHeight存在大量冗余操作
    
    空间复杂度：O(N)
    取决于节点数，不会超过n
*/


var isBalanced=function(root){
    if(!root) return true
    let queue=[root],nodes=[]
    while(queue.length){
        let curNode=queue.shift()
        //从底部构建一个树
        nodes.unshift(curNode)
        if(curNode.left){
            queue.push(curNode.left)
        }
        if(curNode.right){
            queue.push(curNode.right)
        }
    }
    for(let node of nodes){
        let left=node.left?node.left.val:0
        let right=node.right?node.right.val:0
        if(Math.abs(left-right)>1){
            return false
        }
        //修改节点val,觉得这个思路蛮好玩的
        node.val=Math.max(left,right)+1
    }
    return true
}

/** 
 迭代。核心在于倒着修改节点val为左右子树的深度值。
 1.构建从底部的节点数组。
 2.遍历节点数组，更改节点val为从当前位置到底部最大深度值。如果左右子树深度值相差大于1，返回false

    时间复杂度：O(N)
    取决于节点个数

    空间复杂度：O(N)
    取决于队列节点数
*/


var isBalanced=function(root){
    if(!root) return true
    let dfs=(root)=>{
        if(!root){
            return 1;
        }
        let left=dfs(root.left)
        let right=dfs(root.right)
        //如果左右深度值为0或差大于1，都返回0
        if(!left||!right||Math.abs(left-right)>1){
            return 0
        }
        //没有，则更新深度值了。
       return Math.max(left,right)+1
    }
    
    return dfs(root)
}

/** 
    返回深度值的深搜方法。后序遍历。先看左右是否出局，然后更新根的深度值

    时间复杂度：O(N)
    取决于树的深度

    空间复杂度：O(N)
    取决于队列节点数
*/