// 589. N叉树的前序遍历
// 给定一个 N 叉树，返回其节点值的前序遍历。

// 例如，给定一个 3叉树 :

 



 

// 返回其前序遍历: [1,3,5,6,2,4]。

 

// 说明: 递归法很简单，你可以使用迭代法完成此题吗?


/*
 * @lc app=leetcode.cn id=589 lang=javascript
 *
 * [589] N叉树的前序遍历
 */

// @lc code=start
/**
 * // Definition for a Node.
 * function Node(val, children) {
 *    this.val = val;
 *    this.children = children;
 * };
 */

/**
 * @param {Node} root
 * @return {number[]}
 */
var preorder = function(root) {
    if(!root){
        return []
    }
    let res=[root.val]
    let dfs=(root)=>{
        if(!root){
            return ;
        }
        for(let i=0;i<root.children.length;i++){
            res.push(root.children[i].val)
            dfs(root.children[i])
        }
    }
   dfs(root)
   return res
};

/** 
    递归，这个还是比较顺利的。
    时间复杂度：O(N)
    取决于树的节点个数。
    空间复杂度：O(N)
    取决于节点个数
*/


var preorder = function(root) {
    if(!root){
        return []
    }
    let queue=[root],res=[]
    while(queue.length){
        let curNode=queue.shift()
        res.push(curNode.val)
        if(curNode.children.length){
            queue=curNode.children.concat(queue)
        }
    }
    return res
};
// @lc code=end

/*
    迭代，这里没有做出来，但是看题解，觉得这个思路挺巧妙的。
    如果存在孩子，就把孩子节点放在前面去连接当前队列(左+右)，用这样的方式去扩充队列队伍。
    循环过程：1.先shift根节点
            2.左子节点+右子节点，赋值为当前队列

    时间复杂度：O(N)
    取决于树的节点个数。
    空间复杂度：O(N)
    取决于节点个数
*/