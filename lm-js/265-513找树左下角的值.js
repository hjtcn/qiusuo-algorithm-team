/*
513. 找树左下角的值
给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。

假设二叉树中至少有一个节点。

 

示例 1:



输入: root = [2,1,3]
输出: 1
示例 2:



输入: [1,2,3,4,null,5,6,null,null,7]
输出: 7
 

提示:

二叉树的节点个数的范围是 [1,104]
-231 <= Node.val <= 231 - 1 

*/

/*
    思路：广搜迭代。返回最后一行节点数组的第一个节点
    题解：递归，不断更新最大深度，同时当前节点即为目标值
*/

var findBottomLeftValue = function(root) {
    let queue=[root],arr=[]
    while(queue.length){
        let len=queue.length
        arr=[]
        for(let i=0;i<len;i++){
            let curNode=queue.shift()
            arr.push(curNode.val)
            if(curNode.left){
                queue.push(curNode.left)
            }
            if(curNode.right){
                queue.push(curNode.right)
            }
        }

    }
    return arr[0]
};
/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/

var findBottomLeftValue = function(root) {
    let maxDepth=0,target=null
    let dfs=(root,depth)=>{
        if(!root) return false
        if(depth>maxDepth){
            maxDepth=depth
            target=root.val
        }
        dfs(root.left,depth+1)
        dfs(root.right,depth+1)
    }
    dfs(root,1)
    return target

}
/*
    时间复杂度：O(N)
    空间复杂度：O(1).不考虑递归栈占据的空间
*/

/*
    思考：更新，打标记，多用善用
*/