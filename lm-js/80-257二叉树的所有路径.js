// 257. 二叉树的所有路径
// 给定一个二叉树，返回所有从根节点到叶子节点的路径。

// 说明: 叶子节点是指没有子节点的节点。

// 示例:

// 输入:

//    1
//  /   \
// 2     3
//  \
//   5

// 输出: ["1->2->5", "1->3"]

// 解释: 所有根节点到叶子节点的路径为: 1->2->5, 1->3


/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @return {string[]}
 */
var binaryTreePaths = function(root) {
    let res=[],road=[]
    const helper=(root)=>{
        if(!root){
            return null
        }
        road.push(root.val)
        if(!root.left&&!root.right) {
            res.push([...road])  
        }
        helper(root.left)
        helper(root.right)
        road.pop()
    }
    helper(root)
    for(let i=0;i<res.length;i++){
        res[i]=res[i].join('->')
    }
    return res
};
//递归已经可以很顺利的自己写出来了，棒棒哒
// 时间复杂度：O(N)
//  取决于树的节点
//  空间复杂度：O(N)
//  除了结果树，还有记录路径的road数组

var binaryTreePaths = function(root) {
    if(!root) return []
    let res=[],queue=[root],road=[],queuePath=['']
    while(queue.length>0){
       let len=queue.length
           let curNode=queue.shift()
           let curPath=queuePath.shift()
           if(!curNode.left&&!curNode.right){
               res.push(curPath+''+curNode.val)
           }
           else{
            curNode.left&&queue.push(curNode.left)&&queuePath.push(curPath+''+curNode.val+'->')
            curNode.right&&queue.push(curNode.right)&&queuePath.push(curPath+''+curNode.val+'->')
           }
    }
    return res
};


/**
 * bfs.这个是看题解get的。还感觉蛮巧妙的
 * 以前只以为广搜就是横向遍历的，也堵塞了这道题的思路。
 * 此次不仅用队列遍历节点，同时也记录的路径
 *  时间复杂度：O(N)
    取决于树的节点
     空间复杂度：O(N)
    除了结果树，还有记录路径的queuePath数组，以及遍历节点的queue数组
 */