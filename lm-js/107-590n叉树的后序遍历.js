// 590. N叉树的后序遍历
// 给定一个 N 叉树，返回其节点值的后序遍历。

// 输入：
// [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
// 输出：
// [2,6,14,11,7,3,12,8,4,13,9,10,5,1]
 

// 说明: 递归法很简单，你可以使用迭代法完成此题吗?


/*
 * @lc app=leetcode.cn id=590 lang=javascript
 *
 * [590] N叉树的后序遍历
 */

// @lc code=start
/**
 * // Definition for a Node.
 * function Node(val,children) {
 *    this.val = val;
 *    this.children = children;
 * };
 */

/**
 * @param {Node} root
 * @return {number[]}
 */
var postorder = function(root) {
    if(!root) return []
    let res=[]
    let dfs=(root)=>{
        if(!root){
            return ;
        }
        if(root.children){
            for(let i=0;i<root.children.length;i++){
                dfs(root.children[i])
                res.push(root.children[i].val)
            }
        }
    }
    dfs(root)
    res.push(root.val)
    return res
};
/** 
 递归。递归还是比较简单的，就是从左到右先children后根。
    时间复杂度：O(N)
    取决于树的节点数
    空间复杂度：O(N)
    取决于树的节点数，递归所占用的临时栈
*/


var postorder = function(root) {
    if(!root) return []
    let res=[]
    let queue=[root]
    while(queue.length){
        let curNode=queue.pop()
        res.unshift(curNode.val)
        if(curNode.children){
            let len=curNode.children.length;
            for(let i=0;i<len;i++){
                queue.push(curNode.children[i])
            }
        }
    }
    return res
};
/** 
 迭代。这个框架我写出来了，但是队列的控制方向完全搞反。还是没模拟出后序遍历的顺序
 错误代码：
     let curNode=queue.shift()
        res.unshift(curNode.val)
        if(curNode.children){
            let len=curNode.children.length;
            for(let i=len-1;i>0;i--){
                queue.push(curNode.children[i])
            }
        }
    时间复杂度：O(N)
    取决于树的节点数
    空间复杂度：O(N)
    取决于树的节点数，queue数组所占空间。
*/
// @lc code=end

