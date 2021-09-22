/*
144. 二叉树的前序遍历
给你二叉树的根节点 root ，返回它节点值的 前序 遍历。

 

示例 1：


输入：root = [1,null,2,3]
输出：[1,2,3]
示例 2：

输入：root = []
输出：[]
示例 3：

输入：root = [1]
输出：[1]
示例 4：


输入：root = [1,2]
输出：[1,2]
示例 5：


输入：root = [1,null,2]
输出：[1,2]
 

提示：

树中节点数目在范围 [0, 100] 内
-100 <= Node.val <= 100
 

进阶：递归算法很简单，你可以通过迭代算法完成吗？
*/
var preorderTraversal = function(root) {
    let target=[]
    let dfs=(root)=>{
        if(!root) return;
        target.push(root.val)
        dfs(root.left)
        dfs(root.right)
    }
    dfs(root)
    return target
};

//这种方法真的不好理解，不好记忆，找了一个比较好理解的方法。
var preorderTraversal = function(root) {
    let target=[]
    let stack=[]
    while(stack.length>0||root){
        while(root){
            target.push(root.val)
            stack.push(root)
            root=root.left
        }
        root=stack.pop()
        root=root.right
    }
    return target
};

/*
    根左右  stack先把根root放入stack.
           stack出栈pop的过程中,先左，再右
           那么入栈的过程，先右，再左。
    个人感觉这种比两层while循环要容易理解一些
*/
var preorderTraversal = function(root) {
    if(!root) return[]
    let target=[]
    let stack=[root]
    while(stack.length>0){
        let curNode=stack.pop()
        target.push(curNode.val)
        if(curNode.right){
            stack.push(curNode.right)
        }
        if(curNode.left){
            stack.push(curNode.left)
        }
    }
    return target
};
/*
    时间复杂度：O(N)
    空间复杂度：O(N)
*/

/*
    思路：简单题，多练习，总结框架

*/