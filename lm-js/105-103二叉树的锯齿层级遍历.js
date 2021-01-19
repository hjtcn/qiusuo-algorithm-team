// 103. 二叉树的锯齿形层序遍历
// 给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

// 例如：
// 给定二叉树 [3,9,20,null,null,15,7],

//     3
//    / \
//   9  20
//     /  \
//    15   7
// 返回锯齿形层序遍历如下：

// [
//   [3],
//   [20,9],
//   [15,7]
// ]

/*
 * @lc app=leetcode.cn id=103 lang=javascript
 *
 * [103] 二叉树的锯齿形层次遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number[][]}
 */
var zigzagLevelOrder = function(root) {
    if(!root) return []
    let queue=[root],res=[],floor=1
    while(queue.length){
        let floorArr=[],len=queue.length
        for(let i=0;i<len;i++){
            let curNode=queue.shift()
            if(floor%2)
            {
                floorArr.push(curNode.val)
            }
            else{
                floorArr.unshift(curNode.val)
            }
            if(curNode.left){
                queue.push(curNode.left)
            }
            if(curNode.right){
                queue.push(curNode.right)
            }
        }
        floor++
        res.push(floorArr)
    }
    return res
};
// @lc code=end


/** 
 BFS.竟然卡了一下，原因是for循环中的队列长度没有提前记录。然后队列已push，for循环就又多了一次。
 多写写，不能放松。
 每一层，看层级是奇偶数，修改层级遍历的方向
 时间复杂度：O(N)
 和树的节点数有关
 空间复杂度：O(N)
 和树的节点数有关。
*/


var zigzagLevelOrder = function(root) {
    if(!root) return []
    let res=[]
    let dfs=(floor,root)=>{
        if(!root) return;
        if(!res[floor]){
            res[floor]=[]
        }
        if(floor%2){
            res[floor].unshift(root.val)
        }
        else{
            res[floor].push(root.val)
        }
        dfs(floor+1,root.left)
        dfs(floor+1,root.right)
    }
    dfs(0,root)
    return res
};

/** 
 递归。递归还是蛮巧妙的
 记录层数和root。目标数组层数为索引值，因此层数从0开始。
 时间复杂度：O(N)
 和树的节点数有关
 空间复杂度：O(N)
 和树的节点数有关。
*/