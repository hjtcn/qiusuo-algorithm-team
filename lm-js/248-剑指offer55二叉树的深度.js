/*
    剑指 Offer 55 - I. 二叉树的深度
    输入一棵二叉树的根节点，求该树的深度。从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。

    例如：

    给定二叉树 [3,9,20,null,null,15,7]，

        3
    / \
    9  20
        /  \
    15   7
    返回它的最大深度 3 。

    

    提示：

    节点总数 <= 10000
*/

/*
    思路：1.递归
          最大深度，返回左右子树的最大深度+1
         2.迭代
          借用队列数组，横向遍历节点，每一行进行count++,最后返回count

*/
var maxDepth = function(root) {
    if(!root) return 0
    let dfs=(root)=>{
        if(!root) return 0
        let left=0,right=0
        if(root.left)
        {
            left=dfs(root.left)+1
        }
        if(root.right){
            right=dfs(root.right)+1
        }
        return Math.max(left,right)
    }
    return dfs(root)+1
};

var maxDepth = function(root) {
    if(!root) return 0
    let queue=[root],count=0
    while(queue.length>0){
        let len=queue.length
        for(let i=0;i<len;i++){
            let curNode=queue.shift()
            if(curNode.left){
                queue.push(curNode.left)
            }
            if(curNode.right){
                queue.push(curNode.right)
            }
        }
        count++;
    }
    return count
};
/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/
/*
    思考：这个题还是比较基础的，加深二叉树遍历的理解。
*/