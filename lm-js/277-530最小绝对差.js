/*
530. 二叉搜索树的最小绝对差
给你一棵所有节点为非负值的二叉搜索树，请你计算树中任意两节点的差的绝对值的最小值。

 

示例：

输入：

   1
    \
     3
    /
   2

输出：
1

解释：
最小绝对差为 1，其中 2 和 1 的差的绝对值为 1（或者 2 和 3）
*/

/*
    思路：中序遍历，记录前一个值，不断更新差值绝对值的最小值
*/
var getMinimumDifference = function(root) {
    let minVal=Infinity,prev=null
    let dfs=(root)=>{
        if(!root) return;
        dfs(root.left)
        if(prev&&minVal>Math.abs(prev.val-root.val)){
            minVal=Math.abs(prev.val-root.val)
        }
        prev=root
        dfs(root.right)
    }
    dfs(root)
    return minVal
};

var getMinimumDifference = function(root) {
    let stack=[],curNode=root,prev=null,minVal=Infinity
    while(curNode||stack.length){
        while(curNode){
            stack.push(curNode)
            curNode=curNode.left
        }
        curNode=stack.pop()
        if(prev&&minVal>Math.abs(prev.val-curNode.val)){
            minVal=Math.abs(prev.val-curNode.val)
        }
        prev=curNode
        curNode=curNode.right
    }
    return minVal
}

/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/

/*
    思考：中序便利迭代还是有些不熟悉的，清是记不住，生气。
*/