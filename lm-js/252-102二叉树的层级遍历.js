/*
    102. 二叉树的层序遍历
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。

 

示例：
二叉树：[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层序遍历结果：

[
  [3],
  [9,20],
  [15,7]
]
*/

/*
    思路：这个是很明显的BFS题，借用数组队列迭代即可。
        想尝试用递归做一下，想了先序遍历，后来发现自己对先序遍历的遍历顺序也不是很清晰，只记住了根左右，后来查询了一下，它更像深搜，根左左一直到没有左节点，然后查当前子树是否有右节点，直到没有右节点，再次回退一步，重复先左后右
    题解：递归，记录深度，通过数组索引
    target[depth]来记录每一层的左右子节点。
*/

var levelOrder = function(root) {
    if(!root) return []
    let queue=[root],target=[]
    while(queue.length>0){
        let len=queue.length
        let arr=[]
        for(let i=0;i<len;i++){
            let curNode=queue.shift()
            curNode.left&&(queue.push(curNode.left))
            curNode.right&&(queue.push(curNode.right))
            arr.push(curNode.val)
        }
        target.push(arr)
    }
    return target
};
// @lc code=end

var levelOrder = function(root) {
    let target=[]
    let dfs=(depth,root)=>
    {
        if(!root) return ;
        target[depth]||(target[depth]=[])
        target[depth].push(root.val)
        depth++
        dfs(depth,root.left)
        dfs(depth,root.right)
    }
    dfs(0,root)
    return target
}


/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/

/*
    思考：这次复习一定要把前中后序遍历搞清楚
*/