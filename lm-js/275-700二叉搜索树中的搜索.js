/*
700. 二叉搜索树中的搜索
给定二叉搜索树（BST）的根节点和一个值。 你需要在BST中找到节点值等于给定值的节点。 返回以该节点为根的子树。 如果节点不存在，则返回 NULL。

例如，

给定二叉搜索树:

        4
       / \
      2   7
     / \
    1   3

和值: 2
你应该返回如下子树:

      2     
     / \   
    1   3
在上述示例中，如果要找的值是 5，但因为没有节点值为 5，我们应该返回 NULL。



*/

/*
    思路：遍历AC之后发现题没看清。二叉搜索树，是可以继续优化的，比如当前节点值小于目标值，那就向节点的右子节点去找。当前值大于目标值，就向节点的左子节点去找。。

    题解：看完题解，觉得自己思维太固定化了。
    迭代就一定要有队列或栈吗？？？
    官方题解也太简洁了。


*/

var searchBST = function(root, val) {
    if(!root) return null
    if(root.val==val) return root
    else if(root.val<val) return searchBST(root.right,val)
    else if(root.val>val) return searchBST(root.left,val) 
};

var searchBST = function(root, val) {
    let stack=[root]
    while(stack.length){
        let curNode=stack.shift()
        if(curNode.val==val){
            return curNode
        }
        else if(curNode.val>val){
            if(curNode.left){
                stack.push(curNode.left)
            }
        }
        else {
            if(curNode.right){
                stack.push(curNode.right)
            }
        }
    }
    return null
}
/*
    时间复杂度：O(logN)
    空间复杂度：O(logN)
*/
var searchBST = function(root, val) {
    while(root&&root.val!=val){
        root=val<root.val?root.left:root.right
    }
    return root
}
/*
    时间复杂度：O(N)
    空间复杂度：O(1)

*/
//普通遍历
var searchBST = function(root, val) {
    if(!root) return null
    if(root.val==val) return root
    return searchBST(root.left,val)||searchBST(root.right,val)   
};
// @lc code=end

var searchBST = function(root, val) {
    let stack=[root]
    while(stack.length){
        let curNode=stack.shift()
        if(curNode.val==val){
            return curNode
        }
        if(curNode.left){
            stack.push(curNode.left)
        }
        if(curNode.right){
            stack.push(curNode.right)
        }
    }
    return null
}

/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/

/*
    思考：不会的时候，要学会自己总结框架。
        稍会的时候，要学会灵活运用框架。
*/