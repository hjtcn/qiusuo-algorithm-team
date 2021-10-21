/*
404. 左叶子之和
计算给定二叉树的所有左叶子之和。

示例：

    3
   / \
  9  20
    /  \
   15   7

在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24
*/

/*
    思路：递归参数中标记上一步它是左节点还是右节点。 dfs(root,flag)
    如果它是叶子节点且还是左叶子节点，则返回当前节点值

    题解：记录父节点，它的左节点是否是叶子结点。

*/

var sumOfLeftLeaves = function(root) {
    if(!root) return 0
    let dfs=(root,flag)=>{
        if(!root) return 0;
        if(!root.left&&!root.right&&flag==1) {
            return root.val
        }
        return dfs(root.left,1)+dfs(root.right,0)
    }
    return dfs(root,0)
};


var sumOfLeftLeaves = function(root) {
    let dfs=(root)=>{
        if(!root) return 0
        let left=dfs(root.left)
        let right=dfs(root.right)
        let rootVal=0
        if(root.left&&!root.left.left&&!root.left.right){
            rootVal=root.left.val
        }
        return left+right+rootVal

    }
    return dfs(root)
}


var sumOfLeftLeaves = function(root) {
    if(!root) return 0
    let queue=[root],sum=0
    while(queue.length){
        let curNode=queue.shift()
        if(curNode.left&&!curNode.left.left&&!curNode.left.right){
            sum+=curNode.left.val
        }
        if(curNode.left) queue.push(curNode.left)
        if(curNode.right) queue.push(curNode.right)
    }
    return sum
}

/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/
/*
    思考：要把二叉树翻来覆去弄明白。无论是从上到下，还是从下至上，还是从两端向中间集中。多试验，见得多了就不怕了。
*/