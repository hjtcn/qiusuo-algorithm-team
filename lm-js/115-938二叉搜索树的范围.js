// 938. 二叉搜索树的范围和
// 给定二叉搜索树的根结点 root，返回值位于范围 [low, high] 之间的所有结点的值的和。

 

// 示例 1：


// 输入：root = [10,5,15,3,7,null,18], low = 7, high = 15
// 输出：32
// 示例 2：


// 输入：root = [10,5,15,3,7,13,18,1,null,6], low = 6, high = 10
// 输出：23
 

// 提示：

// 树中节点数目在范围 [1, 2 * 104] 内
// 1 <= Node.val <= 105
// 1 <= low <= high <= 105
// 所有 Node.val 互不相同

/** 
 思路。就是遍历，如果在区间内的，就相加
 遍历方式：DFS
         BFS
 
都打算commit的时候，我又一想，是二叉搜索树啊，是不是有优化的点呢？
如果当前节点的val值>=low,向左节点找，找最小节点找，直到小于low为止
如果当前节点的val值<=hight,向右节点找，找最大节点，直到大于high为止


*/



/*
 * @lc app=leetcode.cn id=938 lang=javascript
 *
 * [938] 二叉搜索树的范围和
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
 * @param {number} low
 * @param {number} high
 * @return {number}
 */
var rangeSumBST = function(root, low, high) {
    let sum=0
    let dfs=(root)=>{
        if(!root) return ;
        if(root.val>=low&&root.val<=high){
            sum+=root.val
        }
        if(root.val>=low){//二叉搜索树可加的一层判断
            dfs(root.left)
        }
        if(root.val<=high){
            dfs(root.right) 
        }
       
    }
    dfs(root)
    return sum
};
// @lc code=end

var rangeSumBST = function(root, low, high) {
    if(!root) return 0;
    let queue=[root],sum=0
    while(queue.length){
        let curNode=queue.shift()
        if(curNode.val>=low&&curNode.val<=high){
            sum+=curNode.val
        }
        if(curNode.left){
            queue.push(curNode.left)
        }
        if(curNode.right){
            queue.push(curNode.right)
        }
    }
    return sum
}

// 二叉搜索树的迭代进一步优化.跟递归判断基本一样
var rangeSumBST = function(root, low, high) {
    if(!root) return 0;
    let queue=[root],sum=0
    while(queue.length){
        let curNode=queue.shift()
        if(!curNode){
            continue
        }
        if(curNode.val>=low&&curNode.val<=high){
            sum+=curNode.val
        }
        if(curNode.val>=low){
            queue.push(curNode.left)
        }
        if(curNode.val<=high){
            queue.push(curNode.right)
        }
    }
    return sum
}


/** 
 时间复杂度：O(N)
 取决于二叉树节点个数
 空间复杂度：O(N)
 取决于二叉树节点个数
*/